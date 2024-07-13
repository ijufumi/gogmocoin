package rest

// Client ...
type Client interface {
	Ticker
	Status
	OrderBooks
	Trades
}

type client struct {
	ticker
	status
	orderBooks
	trades
}

// New ...
func New() Client {
	return &client{
		ticker:     newTicker(),
		status:     newStatus(),
		orderBooks: newOrderBooks(),
		trades:     newTrades(),
	}
}
