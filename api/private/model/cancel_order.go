package model

import "api_client/api/common/model"

// CancelOrderReq ...
type CancelOrderReq struct {
	OrderID int64 `json:"orderId"`
}

// CancelOrderRes ...
type CancelOrderRes struct {
	model.ResponseCommon
}
