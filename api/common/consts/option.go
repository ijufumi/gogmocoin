package consts

// Option specifies an additional execution constraint on an order.
type Option string

const (
	// OptionTakerOnly rejects an order that would be filled as a maker, so the
	// order only executes when it can be filled immediately as a taker.
	OptionTakerOnly = Option("TAKER_ONLY")
)
