package ws

import (
	"context"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/public/ws/model"
)

type Trades interface {
	Subscribe() error
	SubscribeWithContext(ctx context.Context) error
	Unsubscribe() error
	Receive() <-chan *model.TradesRes
}

type trades struct {
	apiBase *api.WSAPIBase
}

func newTrades(symbol consts.Symbol, option *consts.Option) *trades {
	apiBase := api.NewWSAPIBase(func(command consts.WebSocketCommand) any {
		return model.NewTradesReq(
			command,
			consts.WebSocketChannelTrades,
			symbol,
			option,
		)
	})

	return &trades{
		apiBase: apiBase,
	}
}

func (c *trades) Subscribe() error {
	return c.SubscribeWithContext(context.Background())
}

func (c *trades) SubscribeWithContext(ctx context.Context) error {
	return c.apiBase.Subscribe(ctx)
}

func (c *trades) Unsubscribe() error {
	return c.apiBase.Unsubscribe()
}

func (c *trades) Receive() <-chan *model.TradesRes {
	return api.RetrieveStream[model.TradesRes]("Trades", c.apiBase.Stream())
}
