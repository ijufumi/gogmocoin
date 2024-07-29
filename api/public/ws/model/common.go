package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
)

// ResponseCommon ...
type ResponseCommon struct {
	Channel consts.WebSocketChannel `json:"channel"`
}
