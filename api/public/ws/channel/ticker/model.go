package ticker

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
	Ask       decimal.Decimal `json:"ask"`
	Bid       decimal.Decimal `json:"bid"`
	High      decimal.Decimal `json:"high"`
	Last      decimal.Decimal `json:"last"`
	Low       decimal.Decimal `json:"low"`
	Symbol    consts.Symbol   `json:"symbol"`
	Timestamp time.Time       `json:"timestamp"`
	Volume    decimal.Decimal `json:"volume"`
}
