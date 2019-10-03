package private

import (
	"encoding/json"
	"fmt"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/private/internal/connect"
	"github.com/ijufumi/gogmocoin/api/private/model"
	"github.com/shopspring/decimal"
)

// Order ...
type Order interface {
	Order(symbol configuration.Symbol, side configuration.Side, executionType configuration.ExecutionType, price, size decimal.Decimal) (*model.OrderRes, error)
}

type order struct {
	con *connect.Connection
}

// Order ...
func (c *order) Order(symbol configuration.Symbol, side configuration.Side, executionType configuration.ExecutionType, price, size decimal.Decimal) (*model.OrderRes, error) {
	req := model.OrderReq{
		Symbol:        symbol,
		Side:          side,
		ExecutionType: executionType,
		Size:          size,
	}

	if executionType == configuration.ExecutionTypeLIMIT {
		req.Price = &price
	}

	res, err := c.con.Post(req, "/v1/order")
	if err != nil {
		return nil, err
	}

	orderRes := new(model.OrderRes)
	err = json.Unmarshal(res, orderRes)
	if err != nil {
		return nil, fmt.Errorf("[order]error:%v,body:%s", err, res)
	}

	if len(orderRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", orderRes.Messages)
	}
	return orderRes, nil
}
