package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
)

// AccountAssetsRes ...
type AccountAssetsRes struct {
	Data []struct {
		consts.Symbol  `json:"symbol"`
		Amount         decimal.Decimal `json:"amount"`
		Available      decimal.Decimal `json:"available"`
		ConversionRate decimal.Decimal `json:"conversionRate"`
	} `json:"data"`
	model.ResponseCommon
}
