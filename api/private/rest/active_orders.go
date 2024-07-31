package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
)

// ActiveOrders ...
type ActiveOrders interface {
	ActiveOrders(symbol consts.Symbol, pageNo int) (*model.ActiveOrdersRes, error)
}

func newActiveOrders(apiKey, secretKey string) activeOrders {
	return activeOrders{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type activeOrders struct {
	api.RestAPIBase
}

// ActiveOrders ...
func (c activeOrders) ActiveOrders(symbol consts.Symbol, pageNo int) (*model.ActiveOrdersRes, error) {
	req := url.Values{
		"symbol": {string(symbol)},
		"page":   {fmt.Sprint(pageNo)},
	}
	res, err := c.Get(req, "/v1/activeOrders")
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
