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

// CloseOrder ...
type CloseOrder interface {
	CloseOrder(positionID int64, symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.CloseOrderRes, error)
	CloseOrderWithContext(ctx context.Context, positionID int64, symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.CloseOrderRes, error)
}

func newCloseOrder(apiKey, secretKey string) closeOrder {
	return closeOrder{
		apiBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type closeOrder struct {
	apiBase api.RestAPIBase
}

// CloseOrder ...
func (c *closeOrder) CloseOrder(positionID int64, symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.CloseOrderRes, error) {
	return c.CloseOrderWithContext(context.Background(), positionID, symbol, side, executionType, price, size)
}

// CloseOrderWithContext ...
func (c *closeOrder) CloseOrderWithContext(ctx context.Context, positionID int64, symbol consts.Symbol, side consts.Side, executionType consts.ExecutionType, price, size decimal.Decimal) (*model.CloseOrderRes, error) {
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

	res, err := c.apiBase.Post(ctx, req, "/v1/closeOrder")
	if err != nil {
		return nil, err
	}

	closeOrderRes := new(model.CloseOrderRes)
	err = json.Unmarshal(res, closeOrderRes)
	if err != nil {
		return nil, fmt.Errorf("[CloseOrder]error:%v,body:%s", err, res)
	}

	if !closeOrderRes.Success() {
		return nil, closeOrderRes.Error()
	}
	return closeOrderRes, nil
}
