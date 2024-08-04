package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
)

// OrderBooksRes ...
type OrderBooksRes struct {
	model.ResponseCommon
	Data struct {
		Asks []struct {
			Price decimal.Decimal `json:"price"`
			Size  decimal.Decimal `json:"size"`
		} `json:"asks"`
		Bids []struct {
			Price decimal.Decimal `json:"price"`
			Size  decimal.Decimal `json:"size"`
		} `json:"bids"`
		Symbol consts.Symbol `json:"symbol"`
	} `json:"data"`
}
