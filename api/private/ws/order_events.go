package ws

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/internal"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/model"
)

type OrderEvents interface {
	Subscribe() error
	Unsubscribe() error
	Receive() <-chan *model.OrderEventsRes
}

type orderEvents struct {
	apiBase *internal.PrivateWSAPIBase
}

func newOrderEvents(apiKey, secretKey string, tokenAutomaticExtension bool) *orderEvents {
	return &orderEvents{
		apiBase: internal.NewPrivateWSAPIBase(apiKey, secretKey, tokenAutomaticExtension, func(command consts.WebSocketCommand) any {
			return model.NewOrderEventsReq(
				command,
				consts.WebSocketChannelOrderEvents,
			)
		}),
	}
}

func (e *orderEvents) Subscribe() error {
	return e.apiBase.Subscribe()
}

func (e *orderEvents) Unsubscribe() error {
	return e.apiBase.Unsubscribe()
}

func (e *orderEvents) Receive() <-chan *model.OrderEventsRes {
	return api.RetrieveStream[model.OrderEventsRes]("OrderEvents", e.apiBase.Stream())
}
