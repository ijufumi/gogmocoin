package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
)

// CloseBulkOrderReq is the request of the bulk position close endpoint.
type CloseBulkOrderReq struct {
	Symbol        consts.Symbol        `json:"symbol"`
	Side          consts.Side          `json:"side"`
	ExecutionType consts.ExecutionType `json:"executionType"`
	Price         *decimal.Decimal     `json:"price,omitempty"`
	Size          decimal.Decimal      `json:"size"`
}

// CloseBulkOrderRes is the response of the bulk position close endpoint, carrying the created order ID.
type CloseBulkOrderRes struct {
	model.ResponseCommon
	Data int64 `json:"data,string"`
}
