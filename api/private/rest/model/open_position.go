package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
	"time"
)

// OpenPositionRes ...
type OpenPositionRes struct {
	model.ResponseCommon
	Data struct {
		model.Pagination `json:"pagination"`
		List             []struct {
			PositionID   int64           `json:"positionId"`
			Symbol       consts.Symbol   `json:"symbol"`
			Side         consts.Side     `json:"side"`
			Size         decimal.Decimal `json:"size"`
			OrderdSize   decimal.Decimal `json:"orderdSize"`
			Price        decimal.Decimal `json:"price"`
			LossGain     decimal.Decimal `json:"lossGain"`
			Leverage     decimal.Decimal `json:"leverage"`
			LosscutPrice decimal.Decimal `json:"losscutPrice"`
			Timestamp    time.Time       `json:"timestamp"`
		} `json:"list"`
	}
}
