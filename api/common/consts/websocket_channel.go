package consts

// WebSocketChannel is the name of a GMO Coin WebSocket subscription channel.
type WebSocketChannel string

const (
	// WebSocketChannelTicker is the public channel streaming ticker updates.
	WebSocketChannelTicker = WebSocketChannel("ticker")
	// WebSocketChannelOrderBooks is the public channel streaming order book snapshots.
	WebSocketChannelOrderBooks = WebSocketChannel("orderbooks")
	// WebSocketChannelTrades is the public channel streaming executed trades.
	WebSocketChannelTrades = WebSocketChannel("trades")
	// WebSocketChannelExecutionEvents is the private channel streaming the user's execution events.
	WebSocketChannelExecutionEvents = WebSocketChannel("executionEvents")
	// WebSocketChannelOrderEvents is the private channel streaming the user's order events.
	WebSocketChannelOrderEvents = WebSocketChannel("orderEvents")
	// WebSocketChannelPositionEvents is the private channel streaming the user's position events.
	WebSocketChannelPositionEvents = WebSocketChannel("positionEvents")
	// WebSocketChannelPositionSummaryEvents is the private channel streaming the user's position summary events.
	WebSocketChannelPositionSummaryEvents = WebSocketChannel("positionSummaryEvents")
)
