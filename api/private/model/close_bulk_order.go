package model

import (
	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

// CloseBulkOrderReq...
type CloseBulkOrderReq struct {
	Symbol        configuration.Symbol        `json:"symbol"`
	Side          configuration.Side          `json:"side"`
	ExecutionType configuration.ExecutionType `json:"executionType"`
	Price         *decimal.Decimal            `json:"price,omitempty"`
	Size          decimal.Decimal             `json:"size"`
}

// CloseBulkOrderRes ...
type CloseBulkOrderRes struct {
	model.ResponseCommon
	Data int64 `json:"data,string"`
}
