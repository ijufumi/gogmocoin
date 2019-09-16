package private

import (
	"api_client/api/private/internal/connect"
	"api_client/api/private/model"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Orders ...
type Orders interface {
	Orders(orderID int64) (*model.OrdersRes, error)
}

type orders struct {
	con *connect.Connection
}

func (o *orders) Orders(orderID int64) (*model.OrdersRes, error) {
	param := url.Values{
		"orderId": {strconv.FormatInt(orderID, 10)},
	}

	res, err := o.con.Get(param, "/v1/orders")
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
