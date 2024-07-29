package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
)

// StatusRes ...
type StatusRes struct {
	model.ResponseCommon
	Data struct {
		Status consts.ExchangeStatus `json:"status"`
	}
}
