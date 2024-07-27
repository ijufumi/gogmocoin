package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

// CloseBulkOrderReq...
type CloseBulkOrderReq struct {
	Symbol        consts.Symbol        `json:"symbol"`
	Side          consts.Side          `json:"side"`
	ExecutionType consts.ExecutionType `json:"executionType"`
	Price         *decimal.Decimal     `json:"price,omitempty"`
	Size          decimal.Decimal      `json:"size"`
}

// CloseBulkOrderRes ...
type CloseBulkOrderRes struct {
	model.ResponseCommon
	Data int64 `json:"data,string"`
}
