package configuration

import (
	"slices"
	"strconv"
	"strings"
)

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
	// SymbolDOGE ...
	SymbolDOGE = Symbol("DOGE")
	// SymbolSOL ...
	SymbolSOL = Symbol("SOL")
	// SymbolASTR ...
	SymbolASTR = Symbol("ASTR")
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
	// SymbolDOTJPY ...
	SymbolDOTJPY = Symbol("DOT_JPY")
	// SymbolATOMJPY ...
	SymbolATOMJPY = Symbol("ATOM_JPY")
	// SymbolADAJPY ...
	SymbolADAJPY = Symbol("ADA_JPY")
	// SymbolLINKJPY ...
	SymbolLINKJPY = Symbol("LINK_JPY")
	// SymbolDOGEJPY ...
	SymbolDOGEJPY = Symbol("DOGE_JPY")
	// SymbolSOLJPY ...
	SymbolSOLJPY = Symbol("SOL_JPY")
	// SymbolJPY ...
	SymbolJPY = Symbol("JPY")
	// SymbolFLR ...
	SymbolFLR = Symbol("FLR")
	// SymbolFIL ...
	SymbolFIL = Symbol("FIL")
	// SymbolSAND ...
	SymbolSAND = Symbol("SAND")
	// SymbolCHZ ...
	SymbolCHZ = Symbol("CHZ")
	// SymbolNONE ...
	SymbolNONE = Symbol("")
)

var allSymbols = []Symbol{
	SymbolBTC,
	SymbolETH,
	SymbolBCH,
	SymbolLTC,
	SymbolXRP,
	SymbolXEM,
	SymbolXLM,
	SymbolBAT,
	SymbolOMG,
	SymbolXTZ,
	SymbolQTUM,
	SymbolENJ,
	SymbolDOT,
	SymbolATOM,
	SymbolMKR,
	SymbolDAI,
	SymbolXYM,
	SymbolMONA,
	SymbolFCR,
	SymbolADA,
	SymbolLINK,
	SymbolDOGE,
	SymbolSOL,
	SymbolASTR,
	SymbolBTCJPY,
	SymbolETHJPY,
	SymbolBCHJPY,
	SymbolLTCJPY,
	SymbolXRPJPY,
	SymbolDOTJPY,
	SymbolATOMJPY,
	SymbolADAJPY,
	SymbolLINKJPY,
	SymbolDOGEJPY,
	SymbolSOLJPY,
	SymbolJPY,
	SymbolFLR,
	SymbolFIL,
	SymbolSAND,
	SymbolCHZ,
}

func (c *Symbol) UnmarshalJSON(d []byte) error {
	symbol := Symbol(strings.ReplaceAll(string(d), `"`, ``))
	if !slices.Contains(allSymbols, symbol) {
		return ErrUnsupportedSymbol
	}
	*c = symbol
	return nil
}

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

// IntervalType ...
type IntervalType string

const (
	IntervalType1Min   = IntervalType("1min")
	IntervalType5Min   = IntervalType("5min")
	IntervalType10Min  = IntervalType("10min")
	IntervalType15Min  = IntervalType("15min")
	IntervalType30Min  = IntervalType("30min")
	IntervalType1Hour  = IntervalType("1hour")
	IntervalType4Hour  = IntervalType("4hour")
	IntervalType8Hour  = IntervalType("8hour")
	IntervalType12Hour = IntervalType("12hour")
	IntervalType1Day   = IntervalType("1day")
	IntervalType1Week  = IntervalType("1week")
	IntervalType1Month = IntervalType("1month")
)

const (
	PublicRestAPIHost  = "https://api.coin.z.com/public"
	PrivateRestAPIHost = "https://api.coin.z.com/private"
)

// TierLevel ...
type TierLevel int8

const (
	TierLevel1 = TierLevel(1)
	TierLevel2 = TierLevel(2)
)

var allTierLevels = []TierLevel{
	TierLevel1, TierLevel2,
}

func (c *TierLevel) UnmarshalJSON(d []byte) error {
	n, err := strconv.ParseInt(string(d), 10, 64)
	if err != nil {
		return err
	}
	level := TierLevel(n)
	if !slices.Contains(allTierLevels, level) {
		return ErrUnsupportedTierLevel
	}
	*c = level
	return nil
}
