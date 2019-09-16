package orderbooks

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws/model"

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
	Asks []struct {
		Price decimal.Decimal `json:"price"`
		Size  decimal.Decimal `json:"size"`
	}
	Bids []struct {
		Price decimal.Decimal `json:"price"`
		Size  decimal.Decimal `json:"size"`
	}
	Symbol configuration.Symbol `json:"symbol"`
}
