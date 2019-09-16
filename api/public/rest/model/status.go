package model

import (
	"gogmocoin/api/common/configuration"
	"gogmocoin/api/common/model"
)

// StatusRes ...
type StatusRes struct {
	model.ResponseCommon
	Data struct {
		Status configuration.ExchangeStatus `json:"status"`
	}
}
