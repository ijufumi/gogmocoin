package consts

// OrderStatus is the lifecycle state of an order.
type OrderStatus string

const (
	// OrderStatusWAITING indicates the order is waiting to be activated (e.g. a stop order).
	OrderStatusWAITING = OrderStatus("WAITING")
	// OrderStatusORDERED indicates the order is active on the order book.
	OrderStatusORDERED = OrderStatus("ORDERED")
	// OrderStatusMODIFYING indicates the order is being modified.
	OrderStatusMODIFYING = OrderStatus("MODIFYING")
	// OrderStatusCANCELLING indicates the order is being cancelled.
	OrderStatusCANCELLING = OrderStatus("CANCELLING")
	// OrderStatusCANCELED indicates the order has been cancelled.
	OrderStatusCANCELED = OrderStatus("CANCELED")
	// OrderStatusEXECUTED indicates the order has been fully executed.
	OrderStatusEXECUTED = OrderStatus("EXECUTED")
	// OrderStatusEXPIRED indicates the order has expired without being executed.
	OrderStatusEXPIRED = OrderStatus("EXPIRED")
)
