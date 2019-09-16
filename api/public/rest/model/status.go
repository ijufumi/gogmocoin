package model

import (
	"api_client/api/common/configuration"
	"api_client/api/common/model"
)

// StatusRes ...
type StatusRes struct {
	model.ResponseCommon
	Data struct {
		Status configuration.ExchangeStatus `json:"status"`
	}
}
