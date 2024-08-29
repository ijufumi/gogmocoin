package ws

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
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
	apiBase := api.NewWSAPIBase(func(command consts.WebSocketCommand) any {
		return model.NewTickerReq(
			command,
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
	return api.RetrieveStream[model.TickerRes]("Ticker", c.apiBase.Stream())
}
