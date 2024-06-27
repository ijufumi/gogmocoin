package private

import (
	"os"

	"github.com/ijufumi/gogmocoin/api/private/internal/connect"
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

	AccessToken
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

	accessToken
}

// NewWithKeys create Client instance.
func NewWithKeys(apiKey, secretKey string) Client {
	c := &client{}
	con := connect.New(apiKey, secretKey)
	c.accountMargin.con = con
	c.accountAssets.con = con

	c.orders.con = con
	c.activeOrders.con = con
	c.executions.con = con
	c.lastExecutions.con = con
	c.openPositions.con = con
	c.positionSummary.con = con

	c.order.con = con
	c.changeOrder.con = con
	c.cancelOrder.con = con

	c.closeOrder.con = con
	c.closeBulkOrder.con = con

	c.changeLosscutPrice.con = con

	c.accessToken.con = con

	return c
}

// New ...
func New() Client {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("API_SECRET")
	return NewWithKeys(apiKey, secretKey)
}
