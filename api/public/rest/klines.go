package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/public/rest/model"
	"net/url"
)

type KLines interface {
	KLines(symbol consts.Symbol, intervalType consts.IntervalType, date string) (*model.KLinesRes, error)
	KLinesWithContext(ctx context.Context, symbol consts.Symbol, intervalType consts.IntervalType, date string) (*model.KLinesRes, error)
}

func newKLines() kLines {
	return kLines{
		apiBase: api.NewRestAPIBase(),
	}
}

type kLines struct {
	apiBase api.RestAPIBase
}

// KLines ...
func (k *kLines) KLines(symbol consts.Symbol, intervalType consts.IntervalType, date string) (*model.KLinesRes, error) {
	return k.KLinesWithContext(context.Background(), symbol, intervalType, date)
}

// KLinesWithContext ...
func (k *kLines) KLinesWithContext(ctx context.Context, symbol consts.Symbol, intervalType consts.IntervalType, date string) (*model.KLinesRes, error) {
	param := url.Values{
		"symbol":   {string(symbol)},
		"interval": {string(intervalType)},
		"date":     {date},
	}

	res, err := k.apiBase.Get(ctx, param, "/v1/klines")
	if err != nil {
		return nil, err
	}

	kLinesRes := new(model.KLinesRes)
	err = json.Unmarshal(res, kLinesRes)
	if err != nil {
		return nil, fmt.Errorf("[KLines]error:%v,body:%s", err, res)
	}

	if !kLinesRes.Success() {
		return nil, kLinesRes.Error()
	}

	return kLinesRes, nil
}
