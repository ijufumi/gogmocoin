package model

import (
	"api_client/api/common/configuration"
	"api_client/api/common/model"

	"github.com/shopspring/decimal"
)

// CloseOrderReq ...
type CloseOrderReq struct {
	Symbol         configuration.Symbol        `json:"symbol"`
	Side           configuration.Side          `json:"side"`
	ExecutionType  configuration.ExecutionType `json:"executionType"`
	Price          *decimal.Decimal            `json:"price,omitempty"`
	SettlePosition []SettlePosition            `json:"settlePosition"`
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
