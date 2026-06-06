package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
)

// ChangeOrderReq is the request of the change order endpoint.
type ChangeOrderReq struct {
	OrderID      int64            `json:"orderId"`
	Price        decimal.Decimal  `json:"price"`
	LossCutPrice *decimal.Decimal `json:"losscutPrice,omitempty"`
}

// ChangeOrderRes is the response of the change order endpoint.
type ChangeOrderRes struct {
	model.ResponseCommon
}
