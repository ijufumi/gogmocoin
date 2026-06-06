package consts

// TimeInForce specifies the execution condition of an order (FAK, FAS, FOK or
// SOK) as defined by the GMO Coin API.
type TimeInForce string

const (
	// TimeInForceFAK is Fill And Kill: fill as much as possible immediately and
	// cancel any remaining quantity.
	TimeInForceFAK = TimeInForce("FAK")
	// TimeInForceFAS is Fill And Store: fill as much as possible immediately and
	// keep any remaining quantity on the order book.
	TimeInForceFAS = TimeInForce("FAS")
	// TimeInForceFOK is Fill Or Kill: fill the entire quantity immediately or
	// cancel the whole order.
	TimeInForceFOK = TimeInForce("FOK")
	// TimeInForceSOK is Success Or Kill, used for closing positions in bulk.
	TimeInForceSOK = TimeInForce("SOK")
)
