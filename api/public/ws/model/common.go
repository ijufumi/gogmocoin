package model

import "gogmocoin/api/common/configuration"

// ResponseCommon ...
type ResponseCommon struct {
	Channel configuration.WebSocketChannel `json:"channel"`
}
