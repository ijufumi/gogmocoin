package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
)

// StatusRes is the response of the exchange status endpoint.
type StatusRes struct {
	model.ResponseCommon
	Data struct {
		Status consts.ExchangeStatus `json:"status"`
	} `json:"data"`
}
