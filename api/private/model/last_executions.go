package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"time"

	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

// LastExecutionsRes ...
type LastExecutionsRes struct {
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
		model.Pagination `json:"pagination"`
	} `json:"data"`
	model.ResponseCommon
}
