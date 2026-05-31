package api

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ijufumi/gogmocoin/v2/api/common/configuration"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/functions"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type state string

const (
	stateStopped          = state("Stopped")
	stateConnecting       = state("Connecting")
	stateConnected        = state("Connected")
	stateConnectionClosed = state("ConnectionClosed")
)

const (
	// reconnectInterval is the wait time before retrying a failed WebSocket dial.
	reconnectInterval = time.Second
	// readLimitBytes caps the size of an incoming WebSocket message.
	readLimitBytes = 10240
	// readDeadline is how long a read may block before the connection is
	// considered dead; it is renewed on every pong.
	readDeadline = 120 * time.Second
	// connectWaitInterval / connectWaitMax bound how long sendMessage waits for
	// the connection to be established before giving up.
	connectWaitInterval = 100 * time.Millisecond
	connectWaitMax      = 10 * time.Second
)

type RequestFactoryFunc func(command consts.WebSocketCommand) any

type WSAPIBase struct {
	conn            atomic.Pointer[websocket.Conn]
	state           *atomic.Value
	startMu         sync.Mutex
	ctx             context.Context
	getHostFunc      HostFactoryFunc
	stopFunc        context.CancelFunc
	subscribeFunc   func() error
	unsubscribeFunc func() error
	stream          chan []byte
	msgStream       chan msgRequest
	// streamMu guards typedStream, the per-session memoized typed stream built
	// lazily by RetrieveStreamOnce.
	streamMu    sync.Mutex
	typedStream any
}

type msgRequest struct {
	msg     any
	errChan chan error
}

func NewWSAPIBase(requestFactory RequestFactoryFunc) *WSAPIBase {
	base := &WSAPIBase{
		state:      &atomic.Value{},
		stream:     make(chan []byte, 100),
		msgStream:  make(chan msgRequest, 1),
		getHostFunc: publicHostFactory,
	}
	base.changeStateToStopped()
	base.setRequestFunc(requestFactory)

	return base
}

func (c *WSAPIBase) SetHostFactoryFunc(f HostFactoryFunc) {
	c.getHostFunc = f
}

func (c *WSAPIBase) initContext(ctx context.Context) {
	newCtx, cancelFunc := context.WithCancel(ctx)
	c.ctx = newCtx
	c.stopFunc = cancelFunc
}

// start launches the send/receive goroutines if they are not already running.
// The check-and-start is guarded by startMu so that concurrent Subscribe calls
// cannot both pass the isStopped check and spawn duplicate goroutine pairs.
func (c *WSAPIBase) start(ctx context.Context) {
	c.startMu.Lock()
	defer c.startMu.Unlock()
	if c.isStopped() {
		c.initContext(ctx)
		// Drop any typed stream from a previous session: it is bound to the old
		// (now cancelled) context and its channel is already closed, so the next
		// RetrieveStreamOnce must build a fresh one against the new context.
		c.resetTypedStream()
		// Move out of the stopped state while still holding the lock so that a
		// concurrent Subscribe observes a non-stopped state and does not spawn a
		// second goroutine pair. The receive goroutine transitions to Connected
		// once dialing succeeds.
		c.changeStateToConnecting()
		go c.doSendGoroutine()
		go c.doReceiveGoroutine()
	}
}

func (c *WSAPIBase) setRequestFunc(f RequestFactoryFunc) {
	c.subscribeFunc = c.createSendFunc(f(consts.WebSocketCommandSubscribe))
	c.unsubscribeFunc = c.createSendFunc(f(consts.WebSocketCommandUnsubscribe))
}

// Subscribe ...
func (c *WSAPIBase) Subscribe(ctx context.Context) error {
	c.start(ctx)
	return c.subscribeFunc()
}

// Unsubscribe ...
func (c *WSAPIBase) Unsubscribe() error {
	defer func() {
		c.close()
	}()
	return c.unsubscribeFunc()
}

func (c *WSAPIBase) createSendFunc(msg any) func() error {
	return func() error {
		// errChan is buffered so that doSendGoroutine can always deliver its
		// result without blocking, even if this caller has already returned
		// because the context was cancelled.
		req := msgRequest{
			msg:     msg,
			errChan: make(chan error, 1),
		}
		ctx := c.ctx
		if ctx == nil {
			ctx = context.Background()
		}
		select {
		case c.msgStream <- req:
		case <-ctx.Done():
			return ctx.Err()
		}
		select {
		case err := <-req.errChan:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (c *WSAPIBase) doSendGoroutine() {
	for {
		select {
		case m := <-c.msgStream:
			e := c.sendMessage(m.msg)
			m.errChan <- e
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *WSAPIBase) doReceiveGoroutine() {
	defer func() {
		if conn := c.conn.Load(); conn != nil && c.isConnected() {
			_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			_ = conn.Close()
		}
	}()

	for {
		// Stop reconnecting/reading once the context is cancelled. Without this
		// guard a persistently failing dial would spin the loop at full CPU and
		// the goroutine would never terminate.
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		if !c.isConnected() {
			if err := c.dial(); err != nil {
				log.Println(err)
				select {
				case <-c.ctx.Done():
					return
				case <-time.After(reconnectInterval):
				}
				continue
			}
		}

		conn := c.conn.Load()
		if conn == nil {
			continue
		}
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("[ReadMessage]error:%v\n", err)
			_ = conn.Close()
			c.changeStateToClosed()
			continue
		}

		select {
		case c.stream <- msg:
		case <-c.ctx.Done():
			return
		}
	}
}

// sendMessage waits for the connection to be established and writes msg as JSON.
func (c *WSAPIBase) sendMessage(msg any) error {
	err := c.waitForConnected()
	if err != nil {
		return err
	}

	conn := c.conn.Load()
	if conn == nil {
		return fmt.Errorf("connection is not established")
	}
	err = conn.WriteJSON(msg)
	if err != nil {
		return fmt.Errorf("write error:%w", err)
	}
	if configuration.IsDebug() {
		log.Printf("[sendMessage]msg: %v", functions.EncodeJSON(msg))
	}
	return nil
}

func (c *WSAPIBase) waitForConnected() error {
	deadline := time.Now().Add(connectWaitMax)
	for {
		if c.isConnected() {
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("connection timeout")
		}
		select {
		case <-c.ctx.Done():
			return c.ctx.Err()
		case <-time.After(connectWaitInterval):
		}
	}
}

func (c *WSAPIBase) dial() error {
	log.Println("dial start")
	if old := c.conn.Load(); old != nil {
		c.changeStateToClosed()
		_ = old.Close()
	}

	c.changeStateToConnecting()
	conn, res, err := websocket.DefaultDialer.Dial(c.getHostFunc(), nil)
	if err != nil {
		c.changeStateToClosed()
		if configuration.IsDebug() {
			return fmt.Errorf("dial error: %v, response: %v, host: %v", err, res, c.getHostFunc())
		}
		return fmt.Errorf("dial error: %v, response: %v", err, res)
	}

	conn.SetReadLimit(readLimitBytes)
	if err := conn.SetReadDeadline(time.Now().Add(readDeadline)); err != nil {
		c.changeStateToClosed()
		_ = conn.Close()
		return fmt.Errorf("set read deadline error: %w", err)
	}
	conn.SetPongHandler(func(appData string) error {
		return conn.SetReadDeadline(time.Now().Add(readDeadline))
	})
	c.conn.Store(conn)
	c.changeStateToConnected()

	return nil
}

// Stream ...
func (c *WSAPIBase) Stream() <-chan []byte {
	return c.stream
}

// resetTypedStream clears the memoized typed stream so the next
// RetrieveStreamOnce rebuilds it for the current session.
func (c *WSAPIBase) resetTypedStream() {
	c.streamMu.Lock()
	c.typedStream = nil
	c.streamMu.Unlock()
}

// close is ...
func (c *WSAPIBase) close() {
	if c.stopFunc != nil {
		c.stopFunc()
	}
	c.changeStateToClosed()
}

func (c *WSAPIBase) isStopped() bool {
	v, ok := c.state.Load().(state)

	if !ok {
		c.changeStateToStopped()
		return true
	}

	return v == stateStopped || v == stateConnectionClosed
}

func (c *WSAPIBase) isConnected() bool {
	v, ok := c.state.Load().(state)

	if !ok {
		c.changeStateToStopped()
		return false
	}

	return v == stateConnected
}

func (c *WSAPIBase) changeStateToStopped() {
	c.state.Store(stateStopped)
}

func (c *WSAPIBase) changeStateToClosed() {
	c.state.Store(stateConnectionClosed)
}

func (c *WSAPIBase) changeStateToConnected() {
	c.state.Store(stateConnected)
}

func (c *WSAPIBase) changeStateToConnecting() {
	c.state.Store(stateConnecting)
}

func publicHostFactory() string {
	return consts.PublicWSAPIHost
}
