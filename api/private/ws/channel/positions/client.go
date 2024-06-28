package positions

import (
	"encoding/json"
	"log"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/private/ws/internal/connect"
)

// Client ...
type Client interface {
	Subscribe() error
	Unsubscribe() error
	Receive() <-chan *Response
}

type client struct {
	conn *connect.Connection
}

// New ...
func New() Client {
	conn := connect.New()
	conn.SetSubscribeFunc(func() interface{} {
		return Request{
			Command: configuration.WebSocketCommandSubscribe,
			Channel: configuration.WebSocketChannelPositions,
		}
	})
	conn.SetUnsubscribeFunc(func() interface{} {
		return Request{
			Command: configuration.WebSocketCommandUnsubscribe,
			Channel: configuration.WebSocketChannelPositions,
		}
	})
	c := &client{
		conn: conn,
	}
	return c
}

func (c *client) Subscribe() error {
	return c.conn.Subscribe()
}

func (c *client) Unsubscribe() error {
	return c.conn.Unsubscribe()
}

func (c *client) Receive() <-chan *Response {
	stream := make(chan *Response, 10)
	go func() {
		for v := range c.conn.Stream() {
			if v == nil {
				return
			}
			log.Printf("received:%v", string(v))
			res := new(Response)
			err := json.Unmarshal(v, res)
			if err != nil {
				log.Printf("[ExecutionEvents] unmarshal error:%v", err)
				continue
			}
			stream <- res

		}
	}()
	return stream
}
