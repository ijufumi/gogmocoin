package model

import (
	"api_client/api/common/configuration"
	"api_client/api/common/model"
	"time"

	"github.com/shopspring/decimal"
)

// OrdersRes ...
type OrdersRes struct {
	Data struct {
		List []struct {
			RootOrderID                 int64 `json:"rootOrderId"`
			OrderID                     int64 `json:"orderId"`
			configuration.Symbol        `json:"symbol"`
			configuration.Side          `json:"side"`
			configuration.OrderType     `json:"orderType"`
			configuration.ExecutionType `json:"executionType"`
			configuration.SettleType    `json:"settleType"`
			Size                        decimal.Decimal           `json:"size"`
			ExecutedSize                decimal.Decimal           `json:"executedSize"`
			Price                       decimal.Decimal           `json:"price"`
			Status                      configuration.OrderStatus `json:"status"`
			configuration.TimeInForce   `json:"timeInForce"`
			Timestamp                   time.Time `json:"timestamp"`
		} `json:"list"`
	} `json:"data"`
	model.ResponseCommon
}
