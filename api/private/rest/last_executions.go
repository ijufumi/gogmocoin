package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
	"strconv"
)

// LastExecutions is the client interface for the latest executions endpoint.
type LastExecutions interface {
	LastExecutions(symbol consts.Symbol, page, count int) (*model.LastExecutionsRes, error)
	LastExecutionsWithContext(ctx context.Context, symbol consts.Symbol, page, count int) (*model.LastExecutionsRes, error)
}

func newLastExecutions(apiKey, secretKey string) lastExecutions {
	return lastExecutions{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type lastExecutions struct {
	api.RestAPIBase
}

// LastExecutions retrieves the most recent executions for the given symbol with pagination using a background context.
func (l *lastExecutions) LastExecutions(symbol consts.Symbol, page, count int) (*model.LastExecutionsRes, error) {
	return l.LastExecutionsWithContext(context.Background(), symbol, page, count)
}

// LastExecutionsWithContext retrieves the most recent executions for the given symbol with pagination using the provided context.
func (l *lastExecutions) LastExecutionsWithContext(ctx context.Context, symbol consts.Symbol, page, count int) (*model.LastExecutionsRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	if page > 0 {
		param.Set("page", strconv.Itoa(page))
	}

	if count > 0 {
		param.Set("count", strconv.Itoa(count))
	}

	res, err := l.Get(ctx, param, "/v1/latestExecutions")
	if err != nil {
		return nil, err
	}

	lastExecutionsRes := new(model.LastExecutionsRes)
	err = json.Unmarshal(res, lastExecutionsRes)
	if err != nil {
		return nil, fmt.Errorf("[LastExecutions]error:%v,body:%s", err, res)
	}

	if !lastExecutionsRes.Success() {
		return nil, lastExecutionsRes.Error()
	}

	return lastExecutionsRes, nil

}
