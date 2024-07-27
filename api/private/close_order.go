package private

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/api/common/api"
	"github.com/ijufumi/gogmocoin/api/common/consts"

	"github.com/ijufumi/gogmocoin/api/private/model"
	"github.com/shopspring/decimal"
)

// CloseOrder ...
type CloseOrder interface {
	CloseOrder(positionID int64, symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.CloseOrderRes, error)
}

func newCloseOrder(apiKey, secretKey string) closeOrder {
	return closeOrder{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type closeOrder struct {
	api.RestAPIBase
}

// CloseOrder ...
func (c *closeOrder) CloseOrder(positionID int64, symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.CloseOrderRes, error) {
	req := model.CloseOrderReq{
		Symbol:        symbol,
		Side:          side,
		ExecutionType: executionType,
		SettlePosition: []model.SettlePosition{{
			PositionID: positionID,
			Size:       size,
		}},
	}

	if executionType == consts.ExecutionTypeLIMIT {
		req.Price = &price
	}

	res, err := c.Post(req, "/v1/closeOrder")
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
