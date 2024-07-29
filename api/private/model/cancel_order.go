package model

import "github.com/ijufumi/gogmocoin/v2/api/common/model"

// CancelOrderReq ...
type CancelOrderReq struct {
	OrderID int64 `json:"orderId"`
}

// CancelOrderRes ...
type CancelOrderRes struct {
	model.ResponseCommon
}
