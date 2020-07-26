package trades

import (
	"time"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/public/ws/model"
	"github.com/shopspring/decimal"
)

// Request is request of ticker.
type Request struct {
	Command              configuration.WebSocketCommand `json:"command"`
	Channel              configuration.WebSocketChannel `json:"channel"`
	configuration.Symbol `json:"symbol"`
	Option               *configuration.Option `json:"option"`
}

// Response is response of ticker.
type Response struct {
	model.ResponseCommon
	Price                decimal.Decimal `json:"price"`
	configuration.Side   `json:"side"`
	Size                 decimal.Decimal `json:"size"`
	Timestamp            time.Time       `json:"timestamp"`
	configuration.Symbol `json:"symbol"`
}
