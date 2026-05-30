package model

import (
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"time"
)

// ResponseCommon is embedded in every REST response and carries the API status
// and any messages returned by the GMO Coin API.
type ResponseCommon struct {
	Messages     []map[string]string `json:"messages,omitempty"`
	Status       consts.Status       `json:"status"`
	ResponseTime time.Time           `json:"responsetime"`
}

// APIError represents a non-OK response from the GMO Coin API. It is returned by
// ResponseCommon.Error and can be inspected with errors.As to recover the status
// code and the per-message details.
type APIError struct {
	Status   consts.Status
	Messages []map[string]string
}

// Error implements the error interface.
func (e *APIError) Error() string {
	return fmt.Sprintf("gmocoin api error: status=%d, messages=%v", e.Status, e.Messages)
}

// Error returns an *APIError carrying the status code and messages so callers
// can branch on the error with errors.As.
func (r *ResponseCommon) Error() error {
	return &APIError{Status: r.Status, Messages: r.Messages}
}

// Success reports whether the API returned an OK status.
func (r *ResponseCommon) Success() bool {
	return r.Status.IsOK()
}

// Pagination ...
type Pagination struct {
	CurrentPage int `json:"currentPage"`
	Count       int `json:"count"`
}

// WebsocketRequestCommon is the base body for WebSocket subscribe/unsubscribe
// commands.
type WebsocketRequestCommon struct {
	Command consts.WebSocketCommand `json:"command"`
	Channel consts.WebSocketChannel `json:"channel"`
}

// WebsocketResponseCommon ...
type WebsocketResponseCommon struct {
	Channel consts.WebSocketChannel `json:"channel"`
	Error   string                  `json:"error"`
}

// PrivateWebsocketResponseCommon ...
type PrivateWebsocketResponseCommon struct {
	WebsocketResponseCommon
	MsgType consts.MsgType `json:"msgType"`
}
