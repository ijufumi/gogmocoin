package consts

// Side ...
type Side string

// Opposite returns opposite side.
func (s Side) Opposite() Side {
	if s == SideBUY {
		return SideSELL
	}
	return SideBUY
}

const (
	// SideBUY ...
	SideBUY = Side("BUY")
	// SideSELL ...
	SideSELL = Side("SELL")
)
