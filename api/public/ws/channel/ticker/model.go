package ticker

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws/model"
	"time"

	"github.com/shopspring/decimal"
)

// Request is request of ticker.
type Request struct {
	Command              configuration.WebSocketCommand `json:"command"`
	Channel              configuration.WebSocketChannel `json:"channel"`
	configuration.Symbol `json:"symbol"`
}

// Response is response of ticker.
type Response struct {
	model.ResponseCommon
	Ask       decimal.Decimal      `json:"ask"`
	Bid       decimal.Decimal      `json:"bid"`
	High      decimal.Decimal      `json:"high"`
	Last      decimal.Decimal      `json:"last"`
	Low       decimal.Decimal      `json:"low"`
	Symbol    configuration.Symbol `json:"symbol"`
	Timestamp time.Time            `json:"timestamp"`
	Volume    decimal.Decimal      `json:"volume"`
}
