package model

import (
	"api_client/api/common/configuration"
	"api_client/api/common/model"

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
