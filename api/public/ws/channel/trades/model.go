package trades

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/public/ws/model"
	"github.com/shopspring/decimal"
	"time"
)

// Request is request of ticker.
type Request struct {
	Command       consts.WebSocketCommand `json:"command"`
	Channel       consts.WebSocketChannel `json:"channel"`
	consts.Symbol `json:"symbol"`
	Option        *consts.Option `json:"option,omitempty"`
}

// Response is response of ticker.
type Response struct {
	model.ResponseCommon
	Price         decimal.Decimal `json:"price"`
	consts.Side   `json:"side"`
	Size          decimal.Decimal `json:"size"`
	Timestamp     time.Time       `json:"timestamp"`
	consts.Symbol `json:"symbol"`
}
