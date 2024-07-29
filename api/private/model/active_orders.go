package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
	"time"
)

// ActiveOrdersRes ...
type ActiveOrdersRes struct {
	model.ResponseCommon
	Data struct {
		model.Pagination `json:"pagination"`
		List             []struct {
			RootOrderID   int64                `json:"rootOrderId"`
			OrderID       int64                `json:"orderId"`
			Symbol        consts.Symbol        `json:"symbol"`
			Side          consts.Side          `json:"side"`
			OrderType     consts.OrderType     `json:"orderType"`
			ExecutionType consts.ExecutionType `json:"executionType"`
			SettleType    consts.SettleType    `json:"settleType"`
			Size          decimal.Decimal      `json:"size"`
			ExecutedSize  decimal.Decimal      `json:"executedSize"`
			Price         decimal.Decimal      `json:"price"`
			LossCutPrice  decimal.Decimal      `json:"losscutPrice"`
			Status        consts.OrderStatus   `json:"status"`
			TimeInForce   consts.TimeInForce   `json:"timeInForce"`
			Timestamp     time.Time            `json:"timestamp"`
		} `json:"list"`
	}
}
