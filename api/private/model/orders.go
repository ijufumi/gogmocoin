package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"time"

	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

// OrdersRes ...
type OrdersRes struct {
	Data struct {
		List []struct {
			RootOrderID          int64 `json:"rootOrderId"`
			OrderID              int64 `json:"orderId"`
			consts.Symbol        `json:"symbol"`
			consts.Side          `json:"side"`
			consts.OrderType     `json:"orderType"`
			consts.ExecutionType `json:"executionType"`
			consts.SettleType    `json:"settleType"`
			Size                 decimal.Decimal    `json:"size"`
			ExecutedSize         decimal.Decimal    `json:"executedSize"`
			Price                decimal.Decimal    `json:"price"`
			LossCutPrice         decimal.Decimal    `json:"losscutPrice"`
			Status               consts.OrderStatus `json:"status"`
			consts.CancelType    `json:"cancelType"`
			consts.TimeInForce   `json:"timeInForce"`
			Timestamp            time.Time `json:"timestamp"`
		} `json:"list"`
	} `json:"data"`
	model.ResponseCommon
}
