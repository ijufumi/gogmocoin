package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
	"time"
)

func NewOrderEventsReq(command consts.WebSocketCommand, channel consts.WebSocketChannel) OrderEventsReq {
	return OrderEventsReq{
		WebsocketRequestCommon: model.WebsocketRequestCommon{
			Command: command,
			Channel: channel,
		},
	}
}

type OrderEventsReq struct {
	model.WebsocketRequestCommon
}

type OrderEventsRes struct {
	model.PrivateWebsocketResponseCommon
	OrderID           int64                `json:"orderId"`
	Symbol            consts.Symbol        `json:"symbol"`
	SettleType        consts.SettleType    `json:"settleType"`
	ExecutionType     consts.ExecutionType `json:"executionType"`
	Side              consts.Side          `json:"side"`
	OrderStatus       consts.OrderStatus   `json:"orderStatus"`
	CancelType        consts.CancelType    `json:"cancelType"`
	OrderTimestamp    time.Time            `json:"orderTimestamp"`
	OrderPrice        decimal.Decimal      `json:"orderPrice"`
	OrderSize         decimal.Decimal      `json:"orderSize"`
	OrderExecutedSize decimal.Decimal      `json:"orderExecutedSize"`
	LosscutPrice      decimal.Decimal      `json:"losscutPrice"`
	TimeInForce       consts.TimeInForce   `json:"timeInForce"`
}
