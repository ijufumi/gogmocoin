package ws

import (
	"context"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/internal"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/model"
)

type PositionSummaryEvents interface {
	Subscribe() error
	SubscribeWithContext(ctx context.Context) error
	Unsubscribe() error
	Receive() <-chan *model.PositionSummaryEventsRes
}

type positionSummaryEvents struct {
	apiBase *internal.PrivateWSAPIBase
}

func newPositionSummaryEvents(apiKey, secretKey string, tokenAutomaticExtension, isPeriodic bool) *positionSummaryEvents {
	return &positionSummaryEvents{
		apiBase: internal.NewPrivateWSAPIBase(apiKey, secretKey, tokenAutomaticExtension, func(command consts.WebSocketCommand) any {
			return model.NewPositionSummaryEventsReq(
				command,
				consts.WebSocketChannelPositionSummaryEvents,
				isPeriodic,
			)
		}),
	}
}

func (c *positionSummaryEvents) Subscribe() error {
	return c.SubscribeWithContext(context.Background())
}

func (c *positionSummaryEvents) SubscribeWithContext(ctx context.Context) error {
	return c.apiBase.Subscribe(ctx)
}

func (e *positionSummaryEvents) Unsubscribe() error {
	return e.apiBase.Unsubscribe()
}

func (e *positionSummaryEvents) Receive() <-chan *model.PositionSummaryEventsRes {
	return api.RetrieveStream[model.PositionSummaryEventsRes]("PositionSummaryEvents", e.apiBase.Stream())
}
