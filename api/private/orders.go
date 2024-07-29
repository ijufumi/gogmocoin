package private

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/model"
	"net/url"
	"strconv"
)

// Orders ...
type Orders interface {
	Orders(orderID int64) (*model.OrdersRes, error)
}

func newOrders(apiKey, secretKey string) orders {
	return orders{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type orders struct {
	api.RestAPIBase
}

func (o *orders) Orders(orderID int64) (*model.OrdersRes, error) {
	param := url.Values{
		"orderId": {strconv.FormatInt(orderID, 10)},
	}

	res, err := o.Get(param, "/v1/orders")
	if err != nil {
		return nil, err
	}
	ordersRes := new(model.OrdersRes)
	err = json.Unmarshal(res, ordersRes)
	if err != nil {
		return nil, fmt.Errorf("[Orders]error:%v,body:%s", err, res)
	}

	if len(ordersRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", ordersRes.Messages)
	}
	return ordersRes, nil
}
