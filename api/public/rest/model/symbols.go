package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

type SymbolsRes struct {
	Data []struct {
		Symbol       consts.Symbol   `json:"symbol"`
		MinOrderSize decimal.Decimal `json:"minOrderSize"`
		MaxOrderSize decimal.Decimal `json:"maxOrderSize"`
		SizeStep     decimal.Decimal `json:"sizeStep"`
		TickSize     decimal.Decimal `json:"tickSize"`
		TakerFee     decimal.Decimal `json:"takerFee"`
		MakerFee     decimal.Decimal `json:"makerFee"`
	}
	model.ResponseCommon
}
