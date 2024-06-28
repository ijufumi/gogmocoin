package executions

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

	OrderID     int64 `json:"orderId"`
	ExecutionID int64 `json:"executionId"`
	PositionID  int64 `json:"positionId"`

	Symbol            string `json:"symbol"`
	SettleType        string `json:"settleType"`
	ExecutionType     string `json:"executionType"`
	Side              string `json:"side"`
	ExecutionPrice    string `json:"executionPrice"`
	ExecutionSize     string `json:"executionSize"`
	LossGain          string `json:"lossGain"`
	Fee               string `json:"fee"`
	OrderPrice        string `json:"orderPrice"`
	OrderSize         string `json:"orderSize"`
	OrderExecutedSize string `json:"orderExecutedSize"`
	TimeInForce       string `json:"timeInForce"`
	MsgType           string `json:"msgType"`

	OrderTimestamp     time.Time `json:"orderTimestamp"`
	ExecutionTimestamp time.Time `json:"executionTimestamp"`
}
