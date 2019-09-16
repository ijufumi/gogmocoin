package model

import "api_client/api/common/configuration"

// ResponseCommon ...
type ResponseCommon struct {
	Channel configuration.WebSocketChannel `json:"channel"`
}
