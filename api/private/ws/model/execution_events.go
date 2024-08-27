package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
	"time"
)

func NewExecutionEventsReq(command consts.WebSocketCommand, channel consts.WebSocketChannel) ExecutionEventsReq {
	return ExecutionEventsReq{
		WebsocketRequestCommon: model.WebsocketRequestCommon{
			Command: command,
			Channel: channel,
		},
	}
}

type ExecutionEventsReq struct {
	model.WebsocketRequestCommon
}

type ExecutionEventsRes struct {
	model.WebsocketResponseCommon
	OrderID            int64                `json:"orderId"`
	ExecutionID        int64                `json:"executionId"`
	Symbol             consts.Symbol        `json:"symbol"`
	SettleType         consts.SettleType    `json:"settleType"`
	ExecutionType      consts.ExecutionType `json:"executionType"`
	Side               consts.Side          `json:"side"`
	ExecutionPrice     decimal.Decimal      `json:"executionPrice"`
	ExecutionSize      decimal.Decimal      `json:"executionSize"`
	PositionID         int64                `json:"positionId"`
	OrderTimestamp     time.Time            `json:"orderTimestamp"`
	ExecutionTimestamp time.Time            `json:"executionTimestamp"`
	LossGain           decimal.Decimal      `json:"lossGain"`
	Fee                decimal.Decimal      `json:"fee"`
	OrderPrice         decimal.Decimal      `json:"orderPrice"`
	OrderSize          decimal.Decimal      `json:"orderSize"`
	OrderExecutedSize  decimal.Decimal      `json:"orderExecutedSize"`
	TimeInForce        consts.TimeInForce   `json:"timeInForce"`
	MsgType            consts.MsgType       `json:"msgType"`
}
