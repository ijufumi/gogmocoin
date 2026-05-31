package ws

import (
	"context"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/internal"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/model"
)

type ExecutionEvents interface {
	Subscribe() error
	SubscribeWithContext(ctx context.Context) error
	Unsubscribe() error
	Stream() <-chan *model.ExecutionEventsRes
	// Deprecated: use Stream instead. Receive is kept as an alias for backward
	// compatibility and will be removed in a future major release.
	Receive() <-chan *model.ExecutionEventsRes
}

type executionEvents struct {
	apiBase *internal.PrivateWSAPIBase
}

func newExecutionEvents(apiKey, secretKey string, tokenAutomaticExtension bool) *executionEvents {
	return &executionEvents{
		apiBase: internal.NewPrivateWSAPIBase(apiKey, secretKey, tokenAutomaticExtension, func(command consts.WebSocketCommand) any {
			return model.NewExecutionEventsReq(
				command,
				consts.WebSocketChannelExecutionEvents,
			)
		}),
	}
}

func (c *executionEvents) Subscribe() error {
	return c.SubscribeWithContext(context.Background())
}

func (c *executionEvents) SubscribeWithContext(ctx context.Context) error {
	return c.apiBase.Subscribe(ctx)
}

func (c *executionEvents) Unsubscribe() error {
	return c.apiBase.Unsubscribe()
}

func (c *executionEvents) Stream() <-chan *model.ExecutionEventsRes {
	return api.RetrieveStreamOnce[model.ExecutionEventsRes](c.apiBase.WS(), "ExecutionEvents")
}

// Deprecated: use Stream instead.
func (c *executionEvents) Receive() <-chan *model.ExecutionEventsRes {
	return c.Stream()
}
