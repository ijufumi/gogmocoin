package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
	"strconv"
)

// LastExecutions ...
type LastExecutions interface {
	LastExecutions(symbol consts.Symbol, page, count int) (*model.LastExecutionsRes, error)
}

func newLastExecutions(apiKey, secretKey string) lastExecutions {
	return lastExecutions{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type lastExecutions struct {
	api.RestAPIBase
}

func (l *lastExecutions) LastExecutions(symbol consts.Symbol, page, count int) (*model.LastExecutionsRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	if page > 0 {
		param.Set("page", strconv.Itoa(page))
	}

	if count > 0 {
		param.Set("count", strconv.Itoa(count))
	}

	res, err := l.Get(param, "/v1/latestExecutions")
	if err != nil {
		return nil, err
	}

	lastExecutionsRes := new(model.LastExecutionsRes)
	err = json.Unmarshal(res, lastExecutionsRes)
	if err != nil {
		return nil, fmt.Errorf("[LastExecutions]error:%v,body:%s", err, res)
	}

	if len(lastExecutionsRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", lastExecutionsRes.Messages)
	}

	return lastExecutionsRes, nil

}
