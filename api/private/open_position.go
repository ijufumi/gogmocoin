package private

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/private/model"
	"net/url"
)

// OpenPositions ...
type OpenPositions interface {
	OpenPositions(symbol consts.Symbol, pageNo int) (*model.OpenPositionRes, error)
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
	req := url.Values{
		"symbol": {string(symbol)},
		"page":   {fmt.Sprint(pageNo)},
	}
	res, err := c.Get(req, "/v1/openPositions")
	if err != nil {
		return nil, err
	}
	opensPositionRes := new(model.OpenPositionRes)
	err = json.Unmarshal(res, opensPositionRes)
	if err != nil {
		return nil, fmt.Errorf("[OpenPositions]error:%v,body:%s", err, res)
	}

	if len(opensPositionRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", opensPositionRes.Messages)
	}

	return opensPositionRes, nil
}
