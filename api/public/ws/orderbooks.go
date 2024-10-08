package ws

import (
	"context"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/public/ws/model"
)

type OrderBooks interface {
	Subscribe() error
	SubscribeWithContext(ctx context.Context) error
	Unsubscribe() error
	Receive() <-chan *model.OrderBooksRes
}

type orderBooks struct {
	apiBase *api.WSAPIBase
}

func newOrderBooks(symbol consts.Symbol) *orderBooks {
	apiBase := api.NewWSAPIBase(func(command consts.WebSocketCommand) any {
		return model.NewOrderBooksReq(
			command,
			consts.WebSocketChannelOrderBooks,
			symbol,
		)
	})

	return &orderBooks{
		apiBase: apiBase,
	}
}

func (c *orderBooks) Subscribe() error {
	return c.SubscribeWithContext(context.Background())
}

func (c *orderBooks) SubscribeWithContext(ctx context.Context) error {
	return c.apiBase.Subscribe(ctx)
}

func (c *orderBooks) Unsubscribe() error {
	return c.apiBase.Unsubscribe()
}

func (c *orderBooks) Receive() <-chan *model.OrderBooksRes {
	return api.RetrieveStream[model.OrderBooksRes]("OrderBooks", c.apiBase.Stream())
}
