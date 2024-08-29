package ws

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/internal"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws/model"
)

type ExecutionEvents interface {
	Subscribe() error
	Unsubscribe() error
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

func (e *executionEvents) Subscribe() error {
	return e.apiBase.Subscribe()
}

func (e *executionEvents) Unsubscribe() error {
	return e.apiBase.Unsubscribe()
}

func (e *executionEvents) Receive() <-chan *model.ExecutionEventsRes {
	return api.RetrieveStream[model.ExecutionEventsRes]("ExecutionEvents", e.apiBase.Stream())
}
