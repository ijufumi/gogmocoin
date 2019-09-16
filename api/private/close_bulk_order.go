package private

import (
	"encoding/json"
	"fmt"
	"gogmocoin/api/common/configuration"
	"gogmocoin/api/private/internal/connect"
	"gogmocoin/api/private/model"

	"github.com/shopspring/decimal"
)

// CloseBulkOrder ...
type CloseBulkOrder interface {
	CloseBulkOrder(symbol configuration.Symbol, side configuration.Side, executionType configuration.ExecutionType, price, size decimal.Decimal) (*model.CloseBulkOrderRes, error)
}

type closeBulkOrder struct {
	con *connect.Connection
}

// CloseBulkOrder ...
func (c *closeBulkOrder) CloseBulkOrder(symbol configuration.Symbol, side configuration.Side, executionType configuration.ExecutionType, price, size decimal.Decimal) (*model.CloseBulkOrderRes, error) {
	request := model.CloseBulkOrderReq{
		Symbol:        symbol,
		Side:          side,
		ExecutionType: executionType,
		Size:          size,
	}

	if executionType == configuration.ExecutionTypeLIMIT {
		request.Price = &price
	}

	res, err := c.con.Post(request, "/v1/closeBulkOrder")
	if err != nil {
		return nil, err
	}

	response := new(model.CloseBulkOrderRes)
	err = json.Unmarshal(res, response)
	if err != nil {
		return nil, fmt.Errorf("[CloseBulkOrder]error:%v,body:%s", err, res)
	}

	if len(response.Messages) != 0 {
		return nil, fmt.Errorf("%v", response.Messages)
	}
	return response, nil
}
