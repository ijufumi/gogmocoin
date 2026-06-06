package consts

// SettleType indicates whether an order opens or closes a position.
type SettleType string

const (
	// SettleTypeOPEN opens a new position.
	SettleTypeOPEN = SettleType("OPEN")
	// SettleTypeCLOSE closes an existing position.
	SettleTypeCLOSE = SettleType("CLOSE")
)
