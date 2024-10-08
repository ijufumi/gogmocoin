package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"github.com/shopspring/decimal"
)

// Order ...
type Order interface {
	Order(symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.OrderRes, error)
	OrderWithContext(ctx context.Context, symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.OrderRes, error)
}

func newOrder(apiKey, secretKey string) order {
	return order{
		apiBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type order struct {
	apiBase api.RestAPIBase
}

// Order ...
func (c *order) Order(symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.OrderRes, error) {
	return c.OrderWithContext(context.Background(), symbol, side, executionType, price, size)
}

// OrderWithContext ...
func (c *order) OrderWithContext(ctx context.Context, symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.OrderRes, error) {
	req := model.OrderReq{
		Symbol:        symbol,
		Side:          side,
		ExecutionType: executionType,
		Size:          size,
	}

	if executionType == consts.ExecutionTypeLIMIT {
		req.Price = &price
	}

	res, err := c.apiBase.Post(ctx, req, "/v1/order")
	if err != nil {
		return nil, err
	}

	orderRes := new(model.OrderRes)
	err = json.Unmarshal(res, orderRes)
	if err != nil {
		return nil, fmt.Errorf("[order]error:%v,body:%s", err, res)
	}

	if !orderRes.Success() {
		return nil, orderRes.Error()
	}
	return orderRes, nil
}
