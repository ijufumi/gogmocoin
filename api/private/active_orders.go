package private

import (
	"api_client/api/common/configuration"
	"api_client/api/private/internal/connect"
	"api_client/api/private/model"
	"encoding/json"
	"fmt"
	"net/url"
)

// ActiveOrders ...
type ActiveOrders interface {
	ActiveOrders(symbol configuration.Symbol, pageNo int) (*model.ActiveOrdersRes, error)
}

type activeOrders struct {
	con *connect.Connection
}

// ActiveOrders ...
func (c activeOrders) ActiveOrders(symbol configuration.Symbol, pageNo int) (*model.ActiveOrdersRes, error) {
	req := url.Values{
		"symbol": {string(symbol)},
		"page":   {fmt.Sprint(pageNo)},
	}
	res, err := c.con.Get(req, "/v1/activeOrders")
	if err != nil {
		return nil, err
	}
	activeOrdersRes := new(model.ActiveOrdersRes)
	err = json.Unmarshal(res, activeOrdersRes)
	if err != nil {
		return nil, fmt.Errorf("[ActiveOrders]error:%v,body:%s", err, res)
	}

	if len(activeOrdersRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", activeOrdersRes.Messages)
	}

	return activeOrdersRes, nil
}
