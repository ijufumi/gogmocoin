package consts

// ExecutionType is the execution method of an order.
type ExecutionType string

const (
	// ExecutionTypeMARKET is a market order executed at the best available price.
	ExecutionTypeMARKET = ExecutionType("MARKET")

	// ExecutionTypeLIMIT is a limit order executed at the specified price or better.
	ExecutionTypeLIMIT = ExecutionType("LIMIT")

	// ExecutionTypeSTOP is a stop order triggered when the market reaches the specified price.
	ExecutionTypeSTOP = ExecutionType("STOP")
)
