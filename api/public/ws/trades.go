package ws

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/public/ws/model"
)

type Trades interface {
	Subscribe() error
	Unsubscribe() error
	Receive() <-chan *model.TradesRes
}

type trades struct {
	*api.WSAPIBase
}

func newTrades(symbol consts.Symbol, option *consts.Option) *trades {
	apiBase := api.NewWSAPIBase()
	apiBase.SetSubscribeFunc(func() interface{} {
		return model.NewTradesReq(
			consts.WebSocketCommandSubscribe,
			consts.WebSocketChannelTrades,
			symbol,
			option,
		)
	})
	apiBase.SetUnsubscribeFunc(func() interface{} {
		return model.NewTradesReq(
			consts.WebSocketCommandUnsubscribe,
			consts.WebSocketChannelTrades,
			symbol,
			option,
		)
	})
	return &trades{
		WSAPIBase: apiBase,
	}
}

func (c *trades) Subscribe() error {
	c.WSAPIBase.Start()
	return c.WSAPIBase.Subscribe()
}

func (c *trades) Unsubscribe() error {
	defer func() {
		c.WSAPIBase.Close()
	}()
	return c.WSAPIBase.Unsubscribe()
}

func (c *trades) Receive() <-chan *model.TradesRes {
	return retrieveStream[model.TradesRes]("Trades", c.WSAPIBase.Stream())
}
