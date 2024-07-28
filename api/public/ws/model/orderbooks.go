package model

import (
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"github.com/ijufumi/gogmocoin/api/common/model"
	"time"

	"github.com/shopspring/decimal"
)

// NewOrderBooksReq ...
func NewOrderBooksReq(command consts.WebSocketCommand, channel consts.WebSocketChannel, symbol consts.Symbol) OrderBooksReq {
	return OrderBooksReq{
		WebsocketRequestCommon: model.WebsocketRequestCommon{
			Command: command,
			Channel: channel,
		},
		Symbol: symbol,
	}
}

// OrderBooksReq is request of orderbooks.
type OrderBooksReq struct {
	model.WebsocketRequestCommon
	consts.Symbol `json:"symbol"`
}

// OrderBooksRes is response of orderbooks.
type OrderBooksRes struct {
	model.WebsocketResponseCommon
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
