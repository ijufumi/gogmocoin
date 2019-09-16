package model

import (
	"api_client/api/common/configuration"
	"api_client/api/common/model"
	"time"

	"github.com/shopspring/decimal"
)

// ExecutionsRes ...
type ExecutionsRes struct {
	Data struct {
		List []struct {
			ExecutionID              int64 `json:"executionId"`
			OrderID                  int64 `json:"orderId"`
			configuration.Symbol     `json:"symbol"`
			configuration.Side       `json:"side"`
			configuration.SettleType `json:"settleType"`
			Size                     decimal.Decimal `json:"size"`
			Price                    decimal.Decimal `json:"price"`
			LossGain                 decimal.Decimal `json:"lossGain"`
			Fee                      decimal.Decimal `json:"fee"`
			Timestamp                time.Time       `json:"timestamp"`
		} `json:"list"`
	} `json:"data"`
	model.ResponseCommon
}
