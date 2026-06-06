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
		RestAPIBase: api.NewRestAPIBase(),
	}
}

type kLines struct {
	api.RestAPIBase
}

// KLines retrieves candlestick data for the given symbol, interval and date using a background context.
func (k *kLines) KLines(symbol consts.Symbol, intervalType consts.IntervalType, date string) (*model.KLinesRes, error) {
	return k.KLinesWithContext(context.Background(), symbol, intervalType, date)
}

// KLinesWithContext retrieves candlestick data for the given symbol, interval and date using the provided context.
func (k *kLines) KLinesWithContext(ctx context.Context, symbol consts.Symbol, intervalType consts.IntervalType, date string) (*model.KLinesRes, error) {
	param := url.Values{
		"symbol":   {string(symbol)},
		"interval": {string(intervalType)},
		"date":     {date},
	}

	res, err := k.Get(ctx, param, "/v1/klines")
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
