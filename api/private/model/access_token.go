package model

import "github.com/ijufumi/gogmocoin/api/common/model"

type AccessTokenRes struct {
	model.ResponseCommon
	Data string `json:"data,omitempty"`
}
