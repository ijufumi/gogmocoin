package consts

// CancelType is the reason an order was cancelled.
type CancelType string

const (
	// CancelTypeUser indicates the order was cancelled by the user.
	CancelTypeUser = CancelType("USER")
	// CancelTypePositionLossCut indicates the order was cancelled by the position loss-cut system.
	CancelTypePositionLossCut = CancelType("POSITION_LOSSCUT")
)
