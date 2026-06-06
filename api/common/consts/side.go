package consts

// Side is the direction of an order, either buy or sell.
type Side string

// Opposite returns opposite side.
func (s Side) Opposite() Side {
	if s == SideBUY {
		return SideSELL
	}
	return SideBUY
}

const (
	// SideBUY is the buy side of an order.
	SideBUY = Side("BUY")
	// SideSELL is the sell side of an order.
	SideSELL = Side("SELL")
)
