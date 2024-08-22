package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
)

// ActiveOrders ...
type ActiveOrders interface {
	ActiveOrders(symbol consts.Symbol, pageNo int) (*model.ActiveOrdersRes, error)
	ActiveOrdersWithContext(ctx context.Context, symbol consts.Symbol, pageNo int) (*model.ActiveOrdersRes, error)
}

func newActiveOrders(apiKey, secretKey string) activeOrders {
	return activeOrders{
		apiBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type activeOrders struct {
	apiBase api.RestAPIBase
}

// ActiveOrders ...
func (c activeOrders) ActiveOrders(symbol consts.Symbol, pageNo int) (*model.ActiveOrdersRes, error) {
	return c.ActiveOrdersWithContext(context.Background(), symbol, pageNo)
}

// ActiveOrdersWithContext ...
func (c activeOrders) ActiveOrdersWithContext(ctx context.Context, symbol consts.Symbol, pageNo int) (*model.ActiveOrdersRes, error) {
	req := url.Values{
		"symbol": {string(symbol)},
		"page":   {fmt.Sprint(pageNo)},
	}
	res, err := c.apiBase.Get(ctx, req, "/v1/activeOrders")
	if err != nil {
		return nil, err
	}
	activeOrdersRes := new(model.ActiveOrdersRes)
	err = json.Unmarshal(res, activeOrdersRes)
	if err != nil {
		return nil, fmt.Errorf("[ActiveOrders]error:%v,body:%s", err, res)
	}

	if !activeOrdersRes.Success() {
		return nil, activeOrdersRes.Error()
	}

	return activeOrdersRes, nil
}
