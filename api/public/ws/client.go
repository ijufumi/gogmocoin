package ws

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
)

// NewOrderBooks returns a WebSocket client for the order book channel of the given symbol.
func NewOrderBooks(symbol consts.Symbol) OrderBooks {
	return newOrderBooks(symbol)
}

// NewTicker returns a WebSocket client for the ticker channel of the given symbol.
func NewTicker(symbol consts.Symbol) Ticker {
	return newTicker(symbol)
}

// NewTrades returns a WebSocket client for the trades channel of the given symbol with an optional execution option.
func NewTrades(symbol consts.Symbol, option *consts.Option) Trades {
	return newTrades(symbol, option)
}
