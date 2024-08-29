package ws

import "github.com/ijufumi/gogmocoin/v2/api/common/configuration"

// NewExecutionEventsWithKeys ...
func NewExecutionEventsWithKeys(apiKey, secretKey string, tokenAutomaticExtension bool) ExecutionEvents {
	return newExecutionEvents(apiKey, secretKey, tokenAutomaticExtension)
}

// NewExecutionEvents ...
func NewExecutionEvents(tokenAutomaticExtension bool) ExecutionEvents {
	return newExecutionEvents(configuration.APIKey(), configuration.APISecret(), tokenAutomaticExtension)
}

// NewOrderEventsWithKeys ...
func NewOrderEventsWithKeys(apiKey, secretKey string, tokenAutomaticExtension bool) OrderEvents {
	return newOrderEvents(apiKey, secretKey, tokenAutomaticExtension)
}

// NewOrderEvents ...
func NewOrderEvents(tokenAutomaticExtension bool) OrderEvents {
	return newOrderEvents(configuration.APIKey(), configuration.APISecret(), tokenAutomaticExtension)
}

// NewPositionEventsWithKeys ...
func NewPositionEventsWithKeys(apiKey, secretKey string, tokenAutomaticExtension bool) PositionEvents {
	return newPositionEvents(apiKey, secretKey, tokenAutomaticExtension)
}

// NewPositionEvents ...
func NewPositionEvents(tokenAutomaticExtension bool) PositionEvents {
	return newPositionEvents(configuration.APIKey(), configuration.APISecret(), tokenAutomaticExtension)
}

// NewPositionSummaryEventsWithKeys ...
func NewPositionSummaryEventsWithKeys(apiKey, secretKey string, tokenAutomaticExtension, isPeriodic bool) PositionSummaryEvents {
	return newPositionSummaryEvents(apiKey, secretKey, tokenAutomaticExtension, isPeriodic)
}

// NewPositionSummaryEvents ...
func NewPositionSummaryEvents(tokenAutomaticExtension, isPeriodic bool) PositionSummaryEvents {
	return newPositionSummaryEvents(configuration.APIKey(), configuration.APISecret(), tokenAutomaticExtension, isPeriodic)
}
