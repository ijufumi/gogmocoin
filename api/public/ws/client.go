package ws

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
)

// NewOrderBooks ...
func NewOrderBooks(symbol consts.Symbol) OrderBooks {
	return newOrderBooks(symbol)
}

// NewTicker ...
func NewTicker(symbol consts.Symbol) Ticker {
	return newTicker(symbol)
}

// NewTrades ...
func NewTrades(symbol consts.Symbol, option *consts.Option) Trades {
	return newTrades(symbol, option)
}
