package orderbooks

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"time"

	"github.com/ijufumi/gogmocoin/api/public/ws/model"
	"github.com/shopspring/decimal"
)

// Request is request of ticker.
type Request struct {
	Command       consts.WebSocketCommand `json:"command"`
	Channel       consts.WebSocketChannel `json:"channel"`
	consts.Symbol `json:"symbol"`
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
	Symbol    consts.Symbol `json:"symbol"`
	Timestamp time.Time     `json:"timestamp"`
}
