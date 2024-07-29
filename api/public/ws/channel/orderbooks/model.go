package orderbooks

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
