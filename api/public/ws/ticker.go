package ws

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/public/ws/model"
)

type Ticker interface {
	Subscribe() error
	Unsubscribe() error
	Receive() <-chan *model.TickerRes
}

type ticker struct {
	apiBase *api.WSAPIBase
}

func newTicker(symbol consts.Symbol) *ticker {
	apiBase := api.NewWSAPIBase()
	apiBase.SetSubscribeFunc(func() interface{} {
		return model.NewTickerReq(
			consts.WebSocketCommandSubscribe,
			consts.WebSocketChannelTicker,
			symbol,
		)
	})
	apiBase.SetUnsubscribeFunc(func() interface{} {
		return model.NewTickerReq(
			consts.WebSocketCommandUnsubscribe,
			consts.WebSocketChannelTicker,
			symbol,
		)
	})
	return &ticker{
		apiBase: apiBase,
	}
}

func (c *ticker) Subscribe() error {
	return c.apiBase.Subscribe()
}

func (c *ticker) Unsubscribe() error {
	return c.apiBase.Unsubscribe()
}

func (c *ticker) Receive() <-chan *model.TickerRes {
	return retrieveStream[model.TickerRes]("Ticker", c.apiBase.Stream())
}
