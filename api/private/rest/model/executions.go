package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
	"time"
)

// ExecutionsRes ...
type ExecutionsRes struct {
	Data struct {
		List []struct {
			ExecutionID       int64 `json:"executionId"`
			OrderID           int64 `json:"orderId"`
			consts.Symbol     `json:"symbol"`
			consts.Side       `json:"side"`
			consts.SettleType `json:"settleType"`
			Size              decimal.Decimal `json:"size"`
			Price             decimal.Decimal `json:"price"`
			LossGain          decimal.Decimal `json:"lossGain"`
			Fee               decimal.Decimal `json:"fee"`
			Timestamp         time.Time       `json:"timestamp"`
		} `json:"list"`
	} `json:"data"`
	model.ResponseCommon
}
