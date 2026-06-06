package consts

// OrderType distinguishes between a user-placed order and a system loss-cut order.
type OrderType string

const (
	// OrderTypeNORMAL is an order placed by the user.
	OrderTypeNORMAL = OrderType("NORMAL")
	// OrderTypeLOSSCUT is an order generated automatically by the loss-cut system.
	OrderTypeLOSSCUT = OrderType("LOSSCUT")
)
