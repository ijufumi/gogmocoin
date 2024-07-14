package rest

// Client ...
type Client interface {
	Ticker
	Status
	OrderBooks
	Trades
	Symbols
}

type client struct {
	ticker
	status
	orderBooks
	trades
	symbols
}

// New ...
func New() Client {
	return &client{
		ticker:     newTicker(),
		status:     newStatus(),
		orderBooks: newOrderBooks(),
		trades:     newTrades(),
		symbols:    newSymbols(),
	}
}
