package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
	"time"
)

// NewTradesReq ...
func NewTradesReq(command consts.WebSocketCommand, channel consts.WebSocketChannel, symbol consts.Symbol, option *consts.Option) TradesReq {
	return TradesReq{
		WebsocketRequestCommon: model.WebsocketRequestCommon{
			Command: command,
			Channel: channel,
		},
		Symbol: symbol,
		Option: option,
	}
}

// TradesReq is request of trades.
type TradesReq struct {
	model.WebsocketRequestCommon
	consts.Symbol `json:"symbol"`
	Option        *consts.Option `json:"option,omitempty"`
}

// TradesRes is response of trades.
type TradesRes struct {
	model.ResponseCommon
	Price         decimal.Decimal `json:"price"`
	consts.Side   `json:"side"`
	Size          decimal.Decimal `json:"size"`
	Timestamp     time.Time       `json:"timestamp"`
	consts.Symbol `json:"symbol"`
}
