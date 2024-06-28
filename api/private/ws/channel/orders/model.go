package orders

import (
	"time"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/private/ws/model"
)

// Request is request of ticker.
type Request struct {
	Command configuration.WebSocketCommand `json:"command"`
	Channel configuration.WebSocketChannel `json:"channel"`
}

// Response represents the JSON structure for execution events
type Response struct {
	model.ResponseCommon

	OrderID int64 `json:"orderId"`

	Symbol            string `json:"symbol"`
	SettleType        string `json:"settleType"`
	ExecutionType     string `json:"executionType"`
	Side              string `json:"side"`
	OrderStatus       string `json:"orderStatus"`
	CancelType        string `json:"cancelType"`
	OrderPrice        string `json:"orderPrice"`
	OrderSize         string `json:"orderSize"`
	OrderExecutedSize string `json:"orderExecutedSize"`
	LosscutPrice      string `json:"losscutPrice"`
	TimeInForce       string `json:"timeInForce"`
	MsgType           string `json:"msgType"`

	OrderTimestamp time.Time `json:"orderTimestamp"` // time.Time for proper timestamp handling
}
