package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
)

// ResponseCommon ...
type ResponseCommon struct {
	Channel consts.WebSocketChannel `json:"channel"`
}
