package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"github.com/ijufumi/gogmocoin/api/common/model"
)

// StatusRes ...
type StatusRes struct {
	model.ResponseCommon
	Data struct {
		Status consts.ExchangeStatus `json:"status"`
	}
}
