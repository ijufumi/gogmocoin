package rest

import (
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
}

func newPositionSummary(apiKey, secretKey string) positionSummary {
	return positionSummary{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type positionSummary struct {
	api.RestAPIBase
}

func (p *positionSummary) PositionSummary(symbol consts.Symbol) (*model.PositionSummaryRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	res, err := p.Get(param, "/v1/positionSummary")
	if err != nil {
		return nil, err
	}

	positionSummaryRes := new(model.PositionSummaryRes)
	err = json.Unmarshal(res, positionSummaryRes)
	if err != nil {
		return nil, fmt.Errorf("[PositionSummary]error:%v,body:%s", err, res)
	}

	if len(positionSummaryRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", positionSummaryRes.Messages)
	}

	return positionSummaryRes, nil
}
