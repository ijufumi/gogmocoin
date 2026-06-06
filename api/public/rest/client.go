package rest

// Client is the public REST API client composing every public endpoint.
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

// New returns a public REST API client ready to call every public endpoint.
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
