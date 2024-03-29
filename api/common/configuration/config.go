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

	// ExecutionTypeSTOP ...
	ExecutionTypeSTOP = ExecutionType("STOP")
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
	// SymbolXEM ...
	SymbolXEM = Symbol("XEM")
	// SymbolXLM ...
	SymbolXLM = Symbol("XLM")
	// SymbolBAT ...
	SymbolBAT = Symbol("BAT")
	// SymbolOMG ...
	SymbolOMG = Symbol("OMG")
	// SymbolXTZ ...
	SymbolXTZ = Symbol("XTZ")
	// SymbolQTUM ...
	SymbolQTUM = Symbol("QTUM")
	// SymbolENJ ...
	SymbolENJ = Symbol("ENJ")
	// SymbolDOT ...
	SymbolDOT = Symbol("DOT")
	// SymbolATOM ...
	SymbolATOM = Symbol("ATOM")
	// SymbolMKR ...
	SymbolMKR = Symbol("MKR")
	// SymbolDAI ...
	SymbolDAI = Symbol("DAI")
	// SymbolXYM ...
	SymbolXYM = Symbol("XYM")
	// SymbolMONA ...
	SymbolMONA = Symbol("MONA")
	// SymbolFCR ...
	SymbolFCR = Symbol("FCR")
	// SymbolADA ...
	SymbolADA = Symbol("ADA")
	// SymbolLINK ...
	SymbolLINK = Symbol("LINK")
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
	// WebSocketCommandSubscribe ...
	WebSocketCommandSubscribe = WebSocketCommand("subscribe")
	// WebSocketCommandUnsubscribe ...
	WebSocketCommandUnsubscribe = WebSocketCommand("unsubscribe")
)

// WebSocketChannel ...
type WebSocketChannel string

const (
	// WebSocketChannelTicker ...
	WebSocketChannelTicker = WebSocketChannel("ticker")
	// WebSocketChannelOrderBooks ...
	WebSocketChannelOrderBooks = WebSocketChannel("orderbooks")
	// WebSocketChannelTrades ...
	WebSocketChannelTrades = WebSocketChannel("trades")
)

// ExchangeStatus ...
type ExchangeStatus string

const (
	// ExchangeStatusOpen ...
	ExchangeStatusOpen = ExchangeStatus("OPEN")
	// ExchangeStatusPreOpen ...
	ExchangeStatusPreOpen = ExchangeStatus("PREOPEN")
	// ExchangeStatusMaintenance ...
	ExchangeStatusMaintenance = ExchangeStatus("MAINTENANCE")
)

// Option ...
type Option string

const (
	// OptionTakerOnly ...
	OptionTakerOnly = Option("TAKER_ONLY")
)

// CancelType ...
type CancelType string

const (
	// CancelTypeUser ...
	CancelTypeUser = CancelType("USER")
	// CancelTypePositionLossCut ...
	CancelTypePositionLossCut = CancelType("POSITION_LOSSCUT")
)
