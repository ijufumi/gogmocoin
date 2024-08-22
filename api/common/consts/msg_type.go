package consts

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
