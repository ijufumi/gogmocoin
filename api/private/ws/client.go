package ws

import "github.com/ijufumi/gogmocoin/v2/api/common/configuration"

// NewExecutionEventsWithKeys returns an execution events WebSocket client using the given API key and secret.
func NewExecutionEventsWithKeys(apiKey, secretKey string, tokenAutomaticExtension bool) ExecutionEvents {
	return newExecutionEvents(apiKey, secretKey, tokenAutomaticExtension)
}

// NewExecutionEvents returns an execution events WebSocket client using the API key and secret from configuration.
func NewExecutionEvents(tokenAutomaticExtension bool) ExecutionEvents {
	return newExecutionEvents(configuration.APIKey(), configuration.APISecret(), tokenAutomaticExtension)
}

// NewOrderEventsWithKeys returns an order events WebSocket client using the given API key and secret.
func NewOrderEventsWithKeys(apiKey, secretKey string, tokenAutomaticExtension bool) OrderEvents {
	return newOrderEvents(apiKey, secretKey, tokenAutomaticExtension)
}

// NewOrderEvents returns an order events WebSocket client using the API key and secret from configuration.
func NewOrderEvents(tokenAutomaticExtension bool) OrderEvents {
	return newOrderEvents(configuration.APIKey(), configuration.APISecret(), tokenAutomaticExtension)
}

// NewPositionEventsWithKeys returns a position events WebSocket client using the given API key and secret.
func NewPositionEventsWithKeys(apiKey, secretKey string, tokenAutomaticExtension bool) PositionEvents {
	return newPositionEvents(apiKey, secretKey, tokenAutomaticExtension)
}

// NewPositionEvents returns a position events WebSocket client using the API key and secret from configuration.
func NewPositionEvents(tokenAutomaticExtension bool) PositionEvents {
	return newPositionEvents(configuration.APIKey(), configuration.APISecret(), tokenAutomaticExtension)
}

// NewPositionSummaryEventsWithKeys returns a position summary events WebSocket client using the given API key and secret.
func NewPositionSummaryEventsWithKeys(apiKey, secretKey string, tokenAutomaticExtension, isPeriodic bool) PositionSummaryEvents {
	return newPositionSummaryEvents(apiKey, secretKey, tokenAutomaticExtension, isPeriodic)
}

// NewPositionSummaryEvents returns a position summary events WebSocket client using the API key and secret from configuration.
func NewPositionSummaryEvents(tokenAutomaticExtension, isPeriodic bool) PositionSummaryEvents {
	return newPositionSummaryEvents(configuration.APIKey(), configuration.APISecret(), tokenAutomaticExtension, isPeriodic)
}
