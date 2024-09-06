package ws

import (
	"context"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/internal"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/model"
)

type PositionEvents interface {
	Subscribe() error
	SubscribeWithContext(ctx context.Context) error
	Unsubscribe() error
	Receive() <-chan *model.PositionEventsRes
}

type positionEvents struct {
	apiBase *internal.PrivateWSAPIBase
}

func newPositionEvents(apiKey, secretKey string, tokenAutomaticExtension bool) *positionEvents {
	return &positionEvents{
		apiBase: internal.NewPrivateWSAPIBase(apiKey, secretKey, tokenAutomaticExtension, func(command consts.WebSocketCommand) any {
			return model.NewPositionEventsReq(
				command,
				consts.WebSocketChannelPositionEvents,
			)
		}),
	}
}

func (c *positionEvents) Subscribe() error {
	return c.SubscribeWithContext(context.Background())
}

func (c *positionEvents) SubscribeWithContext(ctx context.Context) error {
	return c.apiBase.Subscribe(ctx)
}

func (e *positionEvents) Unsubscribe() error {
	return e.apiBase.Unsubscribe()
}

func (e *positionEvents) Receive() <-chan *model.PositionEventsRes {
	return api.RetrieveStream[model.PositionEventsRes]("PositionEvents", e.apiBase.Stream())
}
