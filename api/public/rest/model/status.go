package model

import (
	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/common/model"
)

// StatusRes ...
type StatusRes struct {
	model.ResponseCommon
	Data struct {
		Status configuration.ExchangeStatus `json:"status"`
	}
}
