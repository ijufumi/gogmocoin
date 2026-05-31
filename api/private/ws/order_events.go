package ws

import (
	"context"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/internal"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/model"
)

type OrderEvents interface {
	Subscribe() error
	SubscribeWithContext(ctx context.Context) error
	Unsubscribe() error
	Stream() <-chan *model.OrderEventsRes
	// Deprecated: use Stream instead. Receive is kept as an alias for backward
	// compatibility and will be removed in a future major release.
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

func (c *orderEvents) Subscribe() error {
	return c.SubscribeWithContext(context.Background())
}

func (c *orderEvents) SubscribeWithContext(ctx context.Context) error {
	return c.apiBase.Subscribe(ctx)
}

func (c *orderEvents) Unsubscribe() error {
	return c.apiBase.Unsubscribe()
}

func (c *orderEvents) Stream() <-chan *model.OrderEventsRes {
	return api.RetrieveStreamOnce[model.OrderEventsRes](c.apiBase.WS(), "OrderEvents")
}

// Deprecated: use Stream instead.
func (c *orderEvents) Receive() <-chan *model.OrderEventsRes {
	return c.Stream()
}
