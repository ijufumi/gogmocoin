package private

import (
	"api_client/api/common/configuration"
	"api_client/api/private/internal/connect"
	"api_client/api/private/model"
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

// CloseOrder ...
type CloseOrder interface {
	CloseOrder(positionID int64, symbol configuration.Symbol, side configuration.Side, executionType configuration.ExecutionType, price, size decimal.Decimal) (*model.CloseOrderRes, error)
}

type closeOrder struct {
	con *connect.Connection
}

// CloseOrder ...
func (c *closeOrder) CloseOrder(positionID int64, symbol configuration.Symbol, side configuration.Side, executionType configuration.ExecutionType, price, size decimal.Decimal) (*model.CloseOrderRes, error) {
	req := model.CloseOrderReq{
		Symbol:        symbol,
		Side:          side,
		ExecutionType: executionType,
		SettlePosition: []model.SettlePosition{{
			PositionID: positionID,
			Size:       size,
		}},
	}

	if executionType == configuration.ExecutionTypeLIMIT {
		req.Price = &price
	}

	res, err := c.con.Post(req, "/v1/closeOrder")
	if err != nil {
		return nil, err
	}

	closeOrderRes := new(model.CloseOrderRes)
	err = json.Unmarshal(res, closeOrderRes)
	if err != nil {
		return nil, fmt.Errorf("[CloseOrder]error:%v,body:%s", err, res)
	}

	if len(closeOrderRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", closeOrderRes.Messages)
	}
	return closeOrderRes, nil
}
