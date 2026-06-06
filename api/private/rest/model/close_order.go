package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
)

// CloseOrderReq is the request of the position close endpoint.
type CloseOrderReq struct {
	Symbol         consts.Symbol        `json:"symbol"`
	Side           consts.Side          `json:"side"`
	ExecutionType  consts.ExecutionType `json:"executionType"`
	Price          *decimal.Decimal     `json:"price,omitempty"`
	SettlePosition []SettlePosition     `json:"settlePosition"`
}

// SettlePosition specifies a position and the size to settle when closing.
type SettlePosition struct {
	PositionID int64           `json:"positionId"`
	Size       decimal.Decimal `json:"size"`
}

// CloseOrderRes is the response of the position close endpoint, carrying the created order ID.
type CloseOrderRes struct {
	model.ResponseCommon
	Data int64 `json:"data,string"`
}
