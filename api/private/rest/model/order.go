package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
)

// OrderReq is the request of the new order endpoint.
type OrderReq struct {
	Symbol        consts.Symbol        `json:"symbol"`
	Side          consts.Side          `json:"side"`
	ExecutionType consts.ExecutionType `json:"executionType"`
	Price         *decimal.Decimal     `json:"price,omitempty"`
	LossCutPrice  *decimal.Decimal     `json:"losscutPrice,omitempty"`
	Size          decimal.Decimal      `json:"size"`
}

// OrderRes is the response of the new order endpoint, carrying the created order ID.
type OrderRes struct {
	model.ResponseCommon
	Data int64 `json:"data,string"`
}
