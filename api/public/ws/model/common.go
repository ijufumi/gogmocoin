package model

import "github.com/ijufumi/gogmocoin/api/common/configuration"

// ResponseCommon ...
type ResponseCommon struct {
	Channel configuration.WebSocketChannel `json:"channel"`
}
