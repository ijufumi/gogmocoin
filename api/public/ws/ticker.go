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
	*api.WSAPIBase
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
		WSAPIBase: apiBase,
	}
}

func (c *ticker) Subscribe() error {
	return c.WSAPIBase.Subscribe()
}

func (c *ticker) Unsubscribe() error {
	return c.WSAPIBase.Unsubscribe()
}

func (c *ticker) Receive() <-chan *model.TickerRes {
	return retrieveStream[model.TickerRes]("Ticker", c.WSAPIBase.Stream())
}
