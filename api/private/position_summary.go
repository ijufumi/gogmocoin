package private

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/api/common/api"
	"net/url"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/private/model"
)

// PositionSummary ...
type PositionSummary interface {
	PositionSummary(symbol configuration.Symbol) (*model.PositionSummaryRes, error)
}

func newPositionSummary(apiKey, secretKey string) positionSummary {
	return positionSummary{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type positionSummary struct {
	api.RestAPIBase
}

func (p *positionSummary) PositionSummary(symbol configuration.Symbol) (*model.PositionSummaryRes, error) {
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
