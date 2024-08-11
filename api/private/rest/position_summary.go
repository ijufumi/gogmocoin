package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
)

// PositionSummary ...
type PositionSummary interface {
	PositionSummary(symbol consts.Symbol) (*model.PositionSummaryRes, error)
	PositionSummaryWithContext(ctx context.Context, symbol consts.Symbol) (*model.PositionSummaryRes, error)
}

func newPositionSummary(apiKey, secretKey string) positionSummary {
	return positionSummary{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type positionSummary struct {
	api.RestAPIBase
}

// PositionSummary ...
func (p *positionSummary) PositionSummary(symbol consts.Symbol) (*model.PositionSummaryRes, error) {
	return p.PositionSummaryWithContext(context.Background(), symbol)
}

// PositionSummaryWithContext ...
func (p *positionSummary) PositionSummaryWithContext(ctx context.Context, symbol consts.Symbol) (*model.PositionSummaryRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	res, err := p.Get(ctx, param, "/v1/positionSummary")
	if err != nil {
		return nil, err
	}

	positionSummaryRes := new(model.PositionSummaryRes)
	err = json.Unmarshal(res, positionSummaryRes)
	if err != nil {
		return nil, fmt.Errorf("[PositionSummary]error:%v,body:%s", err, res)
	}

	if !positionSummaryRes.Success() {
		return nil, positionSummaryRes.Error()
	}

	return positionSummaryRes, nil
}
