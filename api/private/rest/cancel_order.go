package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
)

// CancelOrder is the client interface for the cancel order endpoint.
type CancelOrder interface {
	CancelOrder(orderID int64) error
	CancelOrderWithContext(ctx context.Context, orderID int64) error
}

func newCancelOrder(apiKey, secretKey string) cancelOrder {
	return cancelOrder{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type cancelOrder struct {
	api.RestAPIBase
}

// CancelOrder cancels the order with the given order ID using a background context.
func (c *cancelOrder) CancelOrder(orderID int64) error {
	return c.CancelOrderWithContext(context.Background(), orderID)
}

// CancelOrderWithContext cancels the order with the given order ID using the provided context.
func (c *cancelOrder) CancelOrderWithContext(ctx context.Context, orderID int64) error {
	req := model.CancelOrderReq{OrderID: orderID}
	res, err := c.Post(ctx, req, "/v1/cancelOrder")
	if err != nil {
		return err
	}

	cancelOrderRes := new(model.CancelOrderRes)
	err = json.Unmarshal(res, cancelOrderRes)
	if err != nil {
		return fmt.Errorf("[CancelOrder]error:%v,body:%s", err, res)
	}

	if !cancelOrderRes.Success() {
		return cancelOrderRes.Error()
	}
	return nil

}
