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

// OpenPositions ...
type OpenPositions interface {
	OpenPositions(symbol consts.Symbol, pageNo int) (*model.OpenPositionRes, error)
	OpenPositionsWithContext(ctx context.Context, symbol consts.Symbol, pageNo int) (*model.OpenPositionRes, error)
}

func newOpenPositions(apiKey, secretKey string) openPositions {
	return openPositions{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type openPositions struct {
	api.RestAPIBase
}

// OpenPositions ...
func (c *openPositions) OpenPositions(symbol consts.Symbol, pageNo int) (*model.OpenPositionRes, error) {
	return c.OpenPositionsWithContext(context.Background(), symbol, pageNo)
}

// OpenPositionsWithContext ...
func (c *openPositions) OpenPositionsWithContext(ctx context.Context, symbol consts.Symbol, pageNo int) (*model.OpenPositionRes, error) {
	req := url.Values{
		"symbol": {string(symbol)},
		"page":   {fmt.Sprint(pageNo)},
	}
	res, err := c.Get(ctx, req, "/v1/openPositions")
	if err != nil {
		return nil, err
	}
	opensPositionRes := new(model.OpenPositionRes)
	err = json.Unmarshal(res, opensPositionRes)
	if err != nil {
		return nil, fmt.Errorf("[OpenPositions]error:%v,body:%s", err, res)
	}

	if !opensPositionRes.Success() {
		return nil, opensPositionRes.Error()
	}

	return opensPositionRes, nil
}
