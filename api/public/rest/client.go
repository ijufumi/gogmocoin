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
	orderbooks
	trades
}

// New ...
func New() Client {
	return &client{}
}
