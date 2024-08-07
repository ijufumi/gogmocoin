package api

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type connectionState string

const (
	connectionStateConnecting = connectionState("Connecting")
	connectionStateConnected  = connectionState("Connected")
	connectionStateClosed     = connectionState("Closed")
)

type WSAPIBase struct {
	needsAuth bool
	apiKey    string
	secretKey string
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

func NewWSAPIBase() *WSAPIBase {
	base := &WSAPIBase{
		state:     &atomic.Value{},
		stream:    make(chan []byte, 100),
		msgStream: make(chan msgRequest, 1),
	}
	base.changeStateToClosed()
	return base
}

func NewPrivateWSAPIBase(apiKey, secretKey string) WSAPIBase {
	return WSAPIBase{
		needsAuth: true,
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}
func (c *WSAPIBase) initContext() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	c.ctx = ctx
	c.stopFunc = cancelFunc
}

// Start ...
func (c *WSAPIBase) Start() {
	c.initContext()
	go c.doSendGoroutine()
	go c.doReceiveGoroutine()
}

// SetSubscribeFunc ...
func (c *WSAPIBase) SetSubscribeFunc(f func() interface{}) {
	c.subscribeFunc = c.createSendFunc(f)
}

// SetUnsubscribeFunc ...
func (c *WSAPIBase) SetUnsubscribeFunc(f func() interface{}) {
	c.unsubscribeFunc = c.createSendFunc(f)
}

// Subscribe ...
func (c *WSAPIBase) Subscribe() error {
	return c.subscribeFunc()
}

// Unsubscribe ...
func (c *WSAPIBase) Unsubscribe() error {
	return c.unsubscribeFunc()
}

func (c *WSAPIBase) createSendFunc(f func() interface{}) func() error {
	return func() error {
		req := msgRequest{
			msg:     f(),
			errChan: make(chan error),
		}
		c.msgStream <- req
		return <-req.errChan
	}
}

func (c *WSAPIBase) doSendGoroutine() {
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

func (c *WSAPIBase) doReceiveGoroutine() {
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
				log.Printf("[Subscribe]error:%v\n", e)
				_ = c.conn.Close()
				c.changeStateToClosed()
				continue // TODO:review
			}
		}
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("[ReadMessage]error:%v\n", err)
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

func (c *WSAPIBase) isConnected() bool {
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
func (c *WSAPIBase) Send(msg interface{}) error {
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

func (c *WSAPIBase) waitForConnected() error {
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

func (c *WSAPIBase) dial() error {
	c.Lock()
	defer c.Unlock()

	log.Println("dial start")
	if c.conn != nil {
		c.changeStateToClosed()
		_ = c.conn.Close()
	}

	c.changeStateToConnecting()
	conn, res, err := websocket.DefaultDialer.Dial(c.getHost(), nil)
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
func (c *WSAPIBase) Stream() <-chan []byte {
	return c.stream
}

// Close is ...
func (c *WSAPIBase) Close() {
	if c.stopFunc != nil {
		c.stopFunc()
	}
}

func (c *WSAPIBase) changeStateToClosed() {
	c.state.Store(connectionStateClosed)
}

func (c *WSAPIBase) changeStateToConnected() {
	c.state.Store(connectionStateConnected)
}

func (c *WSAPIBase) changeStateToConnecting() {
	c.state.Store(connectionStateConnecting)
}

//func (c *WSAPIBase) makeHeader(systemDatetime time.Time, r *http.Request, method, path, body string) {
//	timeStamp := systemDatetime.Unix()*1000 + int64(systemDatetime.Nanosecond())/int64(time.Millisecond)
//	r.Header.Set("API-TIMESTAMP", fmt.Sprint(timeStamp))
//	r.Header.Set("API-KEY", c.apiKey)
//	r.Header.Set("API-SIGN", c.makeSign(c.secretKey, timeStamp, method, path, body))
//}
//
//func (c *WSAPIBase) makeSign(secretKey string, timeStamp int64, method, path, body string) string {
//	h := hmac.New(sha256.New, []byte(secretKey))
//	h.Write([]byte(fmt.Sprintf("%v%v%v%v", timeStamp, method, path, body)))
//	return hex.EncodeToString(h.Sum(nil))
//}

func (c *WSAPIBase) getHost() string {
	if c.needsAuth {
		return consts.PrivateWSAPIHost
	}
	return consts.PublicWSAPIHost
}
