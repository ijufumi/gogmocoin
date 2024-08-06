package model

import "github.com/ijufumi/gogmocoin/v2/api/common/model"

type WSAuthReq struct {
	Token string `json:"token"`
}

type WSAuthRes struct {
	model.ResponseCommon
	Data string `json:"data"`
}
