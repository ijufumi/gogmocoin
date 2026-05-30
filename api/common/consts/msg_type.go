package consts

// MsgType identifies the kind of message delivered on a private WebSocket
// channel (order, execution and position events). The string values mirror the
// codes defined by the GMO Coin API documentation.
type MsgType string

const (
	MsgTypeER       = MsgType("ER")
	MsgTypeNOR      = MsgType("NOR")
	MsgTypeROR      = MsgType("ROR")
	MsgTypeCOR      = MsgType("COR")
	MsgTypeOPR      = MsgType("OPR")
	MsgTypeUPR      = MsgType("UPR")
	MsgTypeULR      = MsgType("ULR")
	MsgTypeCPR      = MsgType("CPR")
	MsgTypeINIT     = MsgType("INIT")
	MsgTypeUPDATE   = MsgType("UPDATE")
	MsgTypePERIODIC = MsgType("PERIODIC")
)
