package configuration

// Debug if you would like to output debug log, set to true.
var Debug = true

// ExecutionType ...
type ExecutionType string

const (
	// ExecutionTypeMARKET ...
	ExecutionTypeMARKET = ExecutionType("MARKET")

	// ExecutionTypeLIMIT ...
	ExecutionTypeLIMIT = ExecutionType("LIMIT")
)

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

// Symbol ...
type Symbol string

const (
	// SymbolBTC ...
	SymbolBTC = Symbol("BTC")
	// SymbolETH ...
	SymbolETH = Symbol("ETH")
	// SymbolBCH ...
	SymbolBCH = Symbol("BCH")
	// SymbolLTC ...
	SymbolLTC = Symbol("LTC")
	// SymbolXRP ...
	SymbolXRP = Symbol("XRP")
	// SymbolBTCJPY ...
	SymbolBTCJPY = Symbol("BTC_JPY")
	// SymbolETHJPY ...
	SymbolETHJPY = Symbol("ETH_JPY")
	// SymbolBCHJPY ...
	SymbolBCHJPY = Symbol("BCH_JPY")
	// SymbolLTCJPY ...
	SymbolLTCJPY = Symbol("LTC_JPY")
	// SymbolXRPJPY ...
	SymbolXRPJPY = Symbol("XRP_JPY")
	// SymbolNONE ...
	SymbolNONE = Symbol("")
)

// OrderType ...
type OrderType string

const (
	// OrderTypeNORMAL ...
	OrderTypeNORMAL = OrderType("NORMAL")
	// OrderTypeLOSSCUT ...
	OrderTypeLOSSCUT = OrderType("LOSSCUT")
)

// SettleType ...
type SettleType string

const (
	// SettleTypeOPEN ...
	SettleTypeOPEN = SettleType("OPEN")
	// SettleTypeCLOSE ...
	SettleTypeCLOSE = SettleType("CLOSE")
)

// OrderStatus ...
type OrderStatus string

const (
	// OrderStatusWAITING ...
	OrderStatusWAITING = OrderStatus("WAITING")
	// OrderStatusORDERED ...
	OrderStatusORDERED = OrderStatus("ORDERED")
	// OrderStatusMODIFYING ...
	OrderStatusMODIFYING = OrderStatus("MODIFYING")
	// OrderStatusCANCELLING ...
	OrderStatusCANCELLING = OrderStatus("CANCELLING")
	// OrderStatusCANCELED ...
	OrderStatusCANCELED = OrderStatus("CANCELED")
	// OrderStatusEXECUTED ...
	OrderStatusEXECUTED = OrderStatus("EXECUTED")
	// OrderStatusEXPIRED ...
	OrderStatusEXPIRED = OrderStatus("EXPIRED")
)

// TimeInForce ...
type TimeInForce string

const (
	// TimeInForceFAK ...
	TimeInForceFAK = TimeInForce("FAK")
	// TimeInForceFAS ...
	TimeInForceFAS = TimeInForce("FAS")
)

// WebSocketCommand ...
type WebSocketCommand string

const (
	WebSocketCommandSubscribe   = WebSocketCommand("subscribe")
	WebSocketCommandUnsubscribe = WebSocketCommand("unsubscribe")
)

// WebSocketChannel ...
type WebSocketChannel string

const (
	WebSocketChannelTicker     = WebSocketChannel("ticker")
	WebSocketChannelOrderBooks = WebSocketChannel("orderbooks")
)

// ExchangeStatus ...
type ExchangeStatus string

const (
	ExchangeStatusOpen        = ExchangeStatus("OPEN")
	ExchangeStatusPreOpen     = ExchangeStatus("PREOPEN")
	ExchangeStatusMaintenance = ExchangeStatus("MAINTENANCE")
)
