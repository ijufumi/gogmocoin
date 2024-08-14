package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"github.com/shopspring/decimal"
)

// ChangeLosscutPrice ...
type ChangeLosscutPrice interface {
	ChangeLosscutPrice(positionID int64, losscutPrice decimal.Decimal) (*model.ChangeLosscutPriceRes, error)
	ChangeLosscutPriceWithContext(ctx context.Context, positionID int64, losscutPrice decimal.Decimal) (*model.ChangeLosscutPriceRes, error)
}

func newChangeLosscutPrice(apiKey, secretKey string) changeLosscutPrice {
	return changeLosscutPrice{
		apiBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type changeLosscutPrice struct {
	apiBase api.RestAPIBase
}

// ChangeLosscutPrice ...
func (c *changeLosscutPrice) ChangeLosscutPrice(positionID int64, losscutPrice decimal.Decimal) (*model.ChangeLosscutPriceRes, error) {
	return c.ChangeLosscutPriceWithContext(context.Background(), positionID, losscutPrice)
}

// ChangeLosscutPriceWithContext ...
func (c *changeLosscutPrice) ChangeLosscutPriceWithContext(ctx context.Context, positionID int64, losscutPrice decimal.Decimal) (*model.ChangeLosscutPriceRes, error) {
	request := model.ChangeLosscutPriceReq{
		PositionID:   positionID,
		LosscutPrice: losscutPrice,
	}

	res, err := c.apiBase.Post(ctx, request, "/v1/changeLosscutPrice")
	if err != nil {
		return nil, err
	}

	response := new(model.ChangeLosscutPriceRes)
	err = json.Unmarshal(res, response)
	if err != nil {
		return nil, fmt.Errorf("[ChangeLosscutPrice]error:%v,body:%s", err, res)
	}

	if !response.Success() {
		return nil, response.Error()
	}
	return response, nil
}
