package model

import (
	"api_client/api/common/configuration"
	"api_client/api/common/model"
	"time"

	"github.com/shopspring/decimal"
)

// TradesRes ...
type TradesRes struct {
	Data struct {
		List []struct {
			Price     decimal.Decimal    `json:"price"`
			Side      configuration.Side `json:"side"`
			Size      decimal.Decimal    `json:"size"`
			Timestamp time.Time          `json:"timestamp"`
		} `json:"list"`
		model.Pagination `json:"pagination"`
	}
	model.ResponseCommon
}
