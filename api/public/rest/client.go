package rest

// Client ...
type Client interface {
	Ticker
	Status
	OrderBooks
	Trades
	KLines
	Symbols
}

type client struct {
	ticker
	status
	orderBooks
	trades
	kLines
	symbols
}

// New ...
func New() Client {
	return &client{
		ticker:     newTicker(),
		status:     newStatus(),
		orderBooks: newOrderBooks(),
		trades:     newTrades(),
		kLines:     newKLines(),
		symbols:    newSymbols(),
	}
}
