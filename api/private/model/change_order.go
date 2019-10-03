package model

import (
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

// ChangeOrderReq ...
type ChangeOrderReq struct {
	OrderID int64           `json:"orderId"`
	Price   decimal.Decimal `json:"price"`
}

// ChangeOrderRes ...
type ChangeOrderRes struct {
	model.ResponseCommon
}
