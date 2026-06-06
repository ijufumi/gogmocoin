package consts

// ExchangeStatus is the operational status of the exchange service.
type ExchangeStatus string

const (
	// ExchangeStatusOpen indicates the exchange is open for trading.
	ExchangeStatusOpen = ExchangeStatus("OPEN")
	// ExchangeStatusPreOpen indicates the exchange is in the pre-open period before trading starts.
	ExchangeStatusPreOpen = ExchangeStatus("PREOPEN")
	// ExchangeStatusMaintenance indicates the exchange is under maintenance and trading is unavailable.
	ExchangeStatusMaintenance = ExchangeStatus("MAINTENANCE")
)
