package consts

// ExecutionType ...
type ExecutionType string

const (
	// ExecutionTypeMARKET ...
	ExecutionTypeMARKET = ExecutionType("MARKET")

	// ExecutionTypeLIMIT ...
	ExecutionTypeLIMIT = ExecutionType("LIMIT")

	// ExecutionTypeSTOP ...
	ExecutionTypeSTOP = ExecutionType("STOP")
)
