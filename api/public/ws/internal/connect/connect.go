package connect

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const host = "wss://api.coin.z.com/ws/public/v1"

type connectionState string

const (
	connectionStateConnecting = connectionState("Connecting")
	connectionStateConnected  = connectionState("Connected")
	connectionStateClosed     = connectionState("Closed")
)

// Connection ...
type Connection struct {
	sync.Mutex
	conn            *websocket.Conn
	state           *atomic.Value
	ctx             context.Context
	stopFunc        context.CancelFunc
	subscribeFunc   func() error
	unsubscribeFunc func() error
	stream          chan []byte
	msgStream       chan msgRequest
}

type msgRequest struct {
	msg     interface{}
	errChan chan error
}

// New is...
func New() *Connection {
	ctx, cancelFunc := context.WithCancel(context.Background())
	conn := &Connection{
		state:     &atomic.Value{},
		ctx:       ctx,
		stopFunc:  cancelFunc,
		stream:    make(chan []byte, 100),
		msgStream: make(chan msgRequest, 1),
	}
	conn.changeStateToClosed()

	go conn.send()
	go conn.receive()
	return conn
}

// SetSubscribeFunc ...
func (c *Connection) SetSubscribeFunc(f func() interface{}) {
	c.subscribeFunc = c.createSendFunc(f)
}

// SetUnsubscribeFunc ...
func (c *Connection) SetUnsubscribeFunc(f func() interface{}) {
	c.unsubscribeFunc = c.createSendFunc(f)
}

// Subscribe ...
func (c *Connection) Subscribe() error {
	return c.subscribeFunc()
}

// Unsubscribe ...
func (c *Connection) Unsubscribe() error {
	return c.unsubscribeFunc()
}

func (c *Connection) createSendFunc(f func() interface{}) func() error {
	return func() error {
		req := msgRequest{
			msg:     f(),
			errChan: make(chan error),
		}
		c.msgStream <- req
		return <-req.errChan
	}
}

func (c *Connection) send() {
	for {
		select {
		case m := <-c.msgStream:
			e := c.Send(m.msg)
			m.errChan <- e
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Connection) receive() {
	defer func() {
		if c.isConnected() {
			_ = c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			_ = c.conn.Close()
		}
	}()

	for {
		if !c.isConnected() {
			if err := c.dial(); err != nil {
				log.Println(err)
				continue
			}
			e := c.subscribeFunc()
			if e != nil {
				log.Println(fmt.Sprintf("[Subscribe]error:%v", e))
				_ = c.conn.Close()
				c.changeStateToClosed()
				continue // TODO:review
			}
		}
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(fmt.Sprintf("[ReadMessage]error:%v", err))
			_ = c.conn.Close()
			c.changeStateToClosed()
			continue // TODO:review
		}

		c.stream <- msg
		select {
		case <-c.ctx.Done():
			return
		default:
		}
	}
}

func (c *Connection) isConnected() bool {
	c.Lock()
	defer c.Unlock()
	v, ok := c.state.Load().(connectionState)

	if !ok {
		c.changeStateToClosed()
		return false
	}

	return v == connectionStateConnected || v == connectionStateConnecting
}

// Send is...
func (c *Connection) Send(msg interface{}) error {
	err := c.waitForConnected()
	if err != nil {
		return err
	}

	err = c.conn.WriteJSON(msg)
	if err != nil {
		return fmt.Errorf("write error:%v", err)
	}
	log.Printf("[Send]msg:%+v", msg)
	return nil
}

func (c *Connection) waitForConnected() error {
	if c.isConnected() {
		return nil
	}

	for i := 0; i < 10; i++ {
		if c.isConnected() {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("connection timeout")
}

func (c *Connection) dial() error {
	c.Lock()
	defer c.Unlock()

	log.Println("dial start")
	if c.conn != nil {
		c.changeStateToClosed()
		_ = c.conn.Close()
	}

	c.changeStateToConnecting()
	conn, res, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		c.changeStateToClosed()
		return fmt.Errorf("dial error:%v, response:%v", err, res)
	}

	conn.SetReadLimit(10240)
	_ = conn.SetReadDeadline(time.Now().Add(120 * time.Second))
	conn.SetPongHandler(func(appData string) error {
		_ = conn.SetReadDeadline(time.Now().Add(120 * time.Second))
		return nil
	})
	c.conn = conn
	c.changeStateToConnected()

	return nil
}

// Stream ...
func (c *Connection) Stream() <-chan []byte {
	return c.stream
}

// Close is ...
func (c *Connection) Close() {
	c.stopFunc()
}

func (c *Connection) changeStateToClosed() {
	c.state.Store(connectionStateClosed)
}

func (c *Connection) changeStateToConnected() {
	c.state.Store(connectionStateConnected)
}

func (c *Connection) changeStateToConnecting() {
	c.state.Store(connectionStateConnecting)
}
