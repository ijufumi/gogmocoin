package model

import "github.com/ijufumi/gogmocoin/v2/api/common/model"

// CancelOrderReq is the request of the cancel order endpoint.
type CancelOrderReq struct {
	OrderID int64 `json:"orderId"`
}

// CancelOrderRes is the response of the cancel order endpoint.
type CancelOrderRes struct {
	model.ResponseCommon
}
