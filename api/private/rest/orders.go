package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
	"strconv"
)

// Orders ...
type Orders interface {
	Orders(orderID int64) (*model.OrdersRes, error)
	OrdersWithContext(ctx context.Context, orderID int64) (*model.OrdersRes, error)
}

func newOrders(apiKey, secretKey string) orders {
	return orders{
		apiBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type orders struct {
	apiBase api.RestAPIBase
}

// Orders ...
func (o *orders) Orders(orderID int64) (*model.OrdersRes, error) {
	return o.OrdersWithContext(context.Background(), orderID)
}

// OrdersWithContext ...
func (o *orders) OrdersWithContext(ctx context.Context, orderID int64) (*model.OrdersRes, error) {
	param := url.Values{
		"orderId": {strconv.FormatInt(orderID, 10)},
	}

	res, err := o.apiBase.Get(ctx, param, "/v1/orders")
	if err != nil {
		return nil, err
	}
	ordersRes := new(model.OrdersRes)
	err = json.Unmarshal(res, ordersRes)
	if err != nil {
		return nil, fmt.Errorf("[Orders]error:%v,body:%s", err, res)
	}

	if !ordersRes.Success() {
		return nil, ordersRes.Error()
	}
	return ordersRes, nil
}
