package model

import (
	"api_client/api/common/configuration"
	"api_client/api/common/model"
	"time"

	"github.com/shopspring/decimal"
)

// ActiveOrdersRes ...
type ActiveOrdersRes struct {
	model.ResponseCommon
	Data struct {
		model.Pagination `json:"pagination"`
		List             []struct {
			RootOrderID   int64                       `json:"rootOrderId"`
			OrderID       int64                       `json:"orderId"`
			Symbol        configuration.Symbol        `json:"symbol"`
			Side          configuration.Side          `json:"side"`
			OrderType     configuration.OrderType     `json:"orderType"`
			ExecutionType configuration.ExecutionType `json:"executionType"`
			SettleType    configuration.SettleType    `json:"settleType"`
			Size          decimal.Decimal             `json:"size"`
			ExecutedSize  decimal.Decimal             `json:"executedSize"`
			Price         decimal.Decimal             `json:"price"`
			Status        configuration.OrderStatus   `json:"status"`
			TimeInForce   configuration.TimeInForce   `json:"timeInForce"`
			Timestamp     time.Time                   `json:"timestamp"`
		} `json:"list"`
	}
}
