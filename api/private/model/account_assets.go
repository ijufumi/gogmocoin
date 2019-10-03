package model

import (
	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

// AccountAssetsRes ...
type AccountAssetsRes struct {
	Data []struct {
		configuration.Symbol `json:"symbol"`
		Amount               decimal.Decimal `json:"amount"`
		Available            decimal.Decimal `json:"available"`
		ConversionRate       decimal.Decimal `json:"conversionRate"`
	} `json:"data"`
	model.ResponseCommon
}
