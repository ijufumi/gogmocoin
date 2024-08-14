package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
)

// CancelOrder ...
type CancelOrder interface {
	CancelOrder(orderID int64) error
	CancelOrderWithContext(ctx context.Context, orderID int64) error
}

func newCancelOrder(apiKey, secretKey string) cancelOrder {
	return cancelOrder{
		apiBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type cancelOrder struct {
	apiBase api.RestAPIBase
}

// CancelOrder ...
func (c *cancelOrder) CancelOrder(orderID int64) error {
	return c.CancelOrderWithContext(context.Background(), orderID)
}

// CancelOrderWithContext ...
func (c *cancelOrder) CancelOrderWithContext(ctx context.Context, orderID int64) error {
	req := model.CancelOrderReq{OrderID: orderID}
	res, err := c.apiBase.Post(ctx, req, "/v1/cancelOrder")
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
