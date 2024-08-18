package ws

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/public/ws/model"
)

type OrderBooks interface {
	Subscribe() error
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
	return c.apiBase.Subscribe()
}

func (c *orderBooks) Unsubscribe() error {
	return c.apiBase.Unsubscribe()
}

func (c *orderBooks) Receive() <-chan *model.OrderBooksRes {
	return retrieveStream[model.OrderBooksRes]("OrderBooks", c.apiBase.Stream())
}
