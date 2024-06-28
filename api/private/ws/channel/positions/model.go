package positions

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

	PositionID  int64     `json:"positionId"`

	Symbol      string    `json:"symbol"`
	Side        string    `json:"side"`
	Size        string    `json:"size"`
	OrderedSize string    `json:"orderdSize"` // Note the apparent typo in "orderdSize"
	Price       string    `json:"price"`
	LossGain    string    `json:"lossGain"`
	Leverage    string    `json:"leverage"`
	LosscutPrice string   `json:"losscutPrice"`
	MsgType     string    `json:"msgType"`

	Timestamp   time.Time `json:"timestamp"` // time.Time for proper timestamp handling
}
