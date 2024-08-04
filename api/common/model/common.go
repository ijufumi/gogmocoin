package model

import (
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"time"
)

// ResponseCommon ...
type ResponseCommon struct {
	Messages     []map[string]string `json:"messages,omitempty"`
	Status       consts.Status       `json:"status"`
	ResponseTime time.Time           `json:"responsetime"`
}

func (r *ResponseCommon) Error() error {
	return fmt.Errorf("%v", r.Messages)
}

func (r *ResponseCommon) Success() bool {
	return r.Status.IsOK()
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
