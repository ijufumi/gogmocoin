package rest

import (
	"os"
)

// Client ...
type Client interface {
	AccountMargin
	AccountAssets

	Orders
	ActiveOrders
	Executions
	LastExecutions
	OpenPositions
	PositionSummary

	Order
	ChangeOrder
	CancelOrder

	CloseOrder
	CloseBulkOrder

	ChangeLosscutPrice

	WSAuth
}

type client struct {
	accountMargin
	accountAssets

	orders
	activeOrders
	executions
	lastExecutions
	openPositions
	positionSummary

	order
	changeOrder
	cancelOrder

	closeOrder
	closeBulkOrder

	changeLosscutPrice

	wsAuth
}

// NewWithKeys create Client instance.
func NewWithKeys(apiKey, secretKey string) Client {
	c := &client{
		accountAssets: newAccountAssets(apiKey, secretKey),
		accountMargin: newAccountMargin(apiKey, secretKey),

		activeOrders:    newActiveOrders(apiKey, secretKey),
		orders:          newOrders(apiKey, secretKey),
		executions:      newExecutions(apiKey, secretKey),
		lastExecutions:  newLastExecutions(apiKey, secretKey),
		openPositions:   newOpenPositions(apiKey, secretKey),
		positionSummary: newPositionSummary(apiKey, secretKey),

		order:       newOrder(apiKey, secretKey),
		changeOrder: newChangeOrder(apiKey, secretKey),
		cancelOrder: newCancelOrder(apiKey, secretKey),

		closeOrder:     newCloseOrder(apiKey, secretKey),
		closeBulkOrder: newCloseBulkOrder(apiKey, secretKey),

		changeLosscutPrice: newChangeLosscutPrice(apiKey, secretKey),

		wsAuth: newWSAuth(apiKey, secretKey),
	}

	return c
}

// New ...
func New() Client {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("API_SECRET")
	return NewWithKeys(apiKey, secretKey)
}
