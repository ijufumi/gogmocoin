package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"time"
)

// ResponseCommon ...
type ResponseCommon struct {
	Messages     []map[string]string `json:"messages,omitempty"`
	Status       int                 `json:"status"`
	ResponseTime time.Time           `json:"responsetime"`
}

// Pagination ...
type Pagination struct {
	CurrentPage int `json:"currentPage"`
	Count       int `json:"count"`
}

type WebsocketRequestCommon struct {
	Command consts.WebSocketCommand `json:"command"`
	Channel consts.WebSocketChannel `json:"channel"`
}

// WebsocketResponseCommon ...
type WebsocketResponseCommon struct {
	Channel consts.WebSocketChannel `json:"channel"`
}
