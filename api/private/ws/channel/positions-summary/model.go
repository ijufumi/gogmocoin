package summary

import (
	"time"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/private/ws/model"
)

// Request is request of ticker.
type Request struct {
	Command configuration.WebSocketCommand `json:"command"`
	Channel configuration.WebSocketChannel `json:"channel"`
	Option  configuration.Option           `json:"option,omitempty"`
}

// Response represents the JSON structure for execution events
type Response struct {
	model.ResponseCommon

	Symbol              string `json:"symbol"`
	Side                string `json:"side"`
	AveragePositionRate string `json:"averagePositionRate"`
	PositionLossGain    string `json:"positionLossGain"`
	SumOrderQuantity    string `json:"sumOrderQuantity"`
	SumPositionQuantity string `json:"sumPositionQuantity"`
	MsgType             string `json:"msgType"`

	Timestamp time.Time `json:"timestamp"` // time.Time for proper timestamp handling
}
