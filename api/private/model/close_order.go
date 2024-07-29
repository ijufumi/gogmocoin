package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
)

// CloseOrderReq ...
type CloseOrderReq struct {
	Symbol         consts.Symbol        `json:"symbol"`
	Side           consts.Side          `json:"side"`
	ExecutionType  consts.ExecutionType `json:"executionType"`
	Price          *decimal.Decimal     `json:"price,omitempty"`
	SettlePosition []SettlePosition     `json:"settlePosition"`
}

// SettlePosition ...
type SettlePosition struct {
	PositionID int64           `json:"positionId"`
	Size       decimal.Decimal `json:"size"`
}

// CloseOrderRes ...
type CloseOrderRes struct {
	model.ResponseCommon
	Data int64 `json:"data,string"`
}
