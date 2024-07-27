package trades

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"time"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/public/ws/model"
	"github.com/shopspring/decimal"
)

// Request is request of ticker.
type Request struct {
	Command       consts.WebSocketCommand        `json:"command"`
	Channel       configuration.WebSocketChannel `json:"channel"`
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
