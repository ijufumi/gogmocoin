package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
	"time"
)

// TradesRes ...
type TradesRes struct {
	Data struct {
		List []struct {
			Price     decimal.Decimal `json:"price"`
			Side      consts.Side     `json:"side"`
			Size      decimal.Decimal `json:"size"`
			Timestamp time.Time       `json:"timestamp"`
		} `json:"list"`
		model.Pagination `json:"pagination"`
	}
	model.ResponseCommon
}
