package consts

// TimeInForce specifies the execution condition of an order (FAK, FAS, FOK or
// SOK) as defined by the GMO Coin API.
type TimeInForce string

const (
	// TimeInForceFAK ...
	TimeInForceFAK = TimeInForce("FAK")
	// TimeInForceFAS ...
	TimeInForceFAS = TimeInForce("FAS")
	// TimeInForceFOK ...
	TimeInForceFOK = TimeInForce("FOK")
	// TimeInForceSOK ...
	TimeInForceSOK = TimeInForce("SOK")
)
