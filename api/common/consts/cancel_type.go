package consts

// CancelType ...
type CancelType string

const (
	// CancelTypeUser ...
	CancelTypeUser = CancelType("USER")
	// CancelTypePositionLossCut ...
	CancelTypePositionLossCut = CancelType("POSITION_LOSSCUT")
)
