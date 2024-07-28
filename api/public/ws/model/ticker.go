package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
	"time"
)

// NewTickerReq ...
func NewTickerReq(command consts.WebSocketCommand, channel consts.WebSocketChannel, symbol consts.Symbol) TickerReq {
	return TickerReq{
		WebsocketRequestCommon: model.WebsocketRequestCommon{
			Command: command,
			Channel: channel,
		},
		Symbol: symbol,
	}
}

// TickerReq is request of ticker.
type TickerReq struct {
	model.WebsocketRequestCommon
	consts.Symbol `json:"symbol"`
}

// TickerRes is response of ticker.
type TickerRes struct {
	model.WebsocketResponseCommon
	Ask       decimal.Decimal `json:"ask"`
	Bid       decimal.Decimal `json:"bid"`
	High      decimal.Decimal `json:"high"`
	Last      decimal.Decimal `json:"last"`
	Low       decimal.Decimal `json:"low"`
	Symbol    consts.Symbol   `json:"symbol"`
	Timestamp time.Time       `json:"timestamp"`
	Volume    decimal.Decimal `json:"volume"`
}
