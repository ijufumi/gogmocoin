package model

import (
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

// AccountMarginRes ...
type AccountMarginRes struct {
	Data struct {
		ProfitLoss       decimal.Decimal `json:"profitLoss"`
		ActualProfitLoss decimal.Decimal `json:"actualProfitLoss"`
		Margin           decimal.Decimal `json:"margin"`
		AvailableAmount  decimal.Decimal `json:"availableAmount"`
	} `json:"data"`
	model.ResponseCommon
}
