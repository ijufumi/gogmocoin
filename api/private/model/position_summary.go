package model

import (
	"api_client/api/common/configuration"
	"api_client/api/common/model"

	"github.com/shopspring/decimal"
)

// PositionSummaryRes ...
type PositionSummaryRes struct {
	Data struct {
		List []struct {
			configuration.Symbol `json:"symbol"`
			configuration.Side   `json:"side"`
			SumPositionQuantity  decimal.Decimal `json:"sumPositionQuantity"`
			SumOrderQuantity     decimal.Decimal `json:"sumOrderQuantity"`
			AveragePositionRate  decimal.Decimal `json:"averagePositionRate"`
			PositionLossGain     decimal.Decimal `json:"positionLossGain"`
		} `json:"list"`
	} `json:"data"`
	model.ResponseCommon
}
