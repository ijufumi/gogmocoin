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
	Stream() <-chan *model.PositionEventsRes
	// Deprecated: use Stream instead. Receive is kept as an alias for backward
	// compatibility and will be removed in a future major release.
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

func (c *positionEvents) Unsubscribe() error {
	return c.apiBase.Unsubscribe()
}

func (c *positionEvents) Stream() <-chan *model.PositionEventsRes {
	return api.RetrieveStreamOnce[model.PositionEventsRes](c.apiBase.WS(), "PositionEvents")
}

// Deprecated: use Stream instead.
func (c *positionEvents) Receive() <-chan *model.PositionEventsRes {
	return c.Stream()
}
