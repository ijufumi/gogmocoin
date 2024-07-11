package trades

import (
	"encoding/json"
	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/public/ws/internal/connect"
	"log"
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
func New(symbol configuration.Symbol, option *configuration.Option) Client {
	conn := connect.New()
	conn.SetSubscribeFunc(func() interface{} {
		return Request{
			Command: configuration.WebSocketCommandSubscribe,
			Channel: configuration.WebSocketChannelTrades,
			Symbol:  symbol,
			Option:  option,
		}
	})
	conn.SetUnsubscribeFunc(func() interface{} {
		return Request{
			Command: configuration.WebSocketCommandUnsubscribe,
			Channel: configuration.WebSocketChannelTrades,
			Symbol:  symbol,
			Option:  option,
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
		for {
			v := <-c.conn.Stream()
			if v == nil {
				return
			}
			log.Printf("received:%v", string(v))
			res := new(Response)
			err := json.Unmarshal(v, res)
			if err != nil {
				log.Printf("[Trades] unmarshal error:%v", err)
				continue
			}
			stream <- res
		}
	}()
	return stream
}
