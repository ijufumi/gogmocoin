package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

// PositionSummaryRes ...
type PositionSummaryRes struct {
	Data struct {
		List []struct {
			consts.Symbol       `json:"symbol"`
			consts.Side         `json:"side"`
			SumPositionQuantity decimal.Decimal `json:"sumPositionQuantity"`
			SumOrderQuantity    decimal.Decimal `json:"sumOrderQuantity"`
			AveragePositionRate decimal.Decimal `json:"averagePositionRate"`
			PositionLossGain    decimal.Decimal `json:"positionLossGain"`
		} `json:"list"`
	} `json:"data"`
	model.ResponseCommon
}
