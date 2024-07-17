package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/api/common/api"
	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/public/rest/model"
	"net/url"
)

type KLines interface {
	KLines(symbol configuration.Symbol, intervalType configuration.IntervalType, date string) (*model.KLinesRes, error)
}

func NewKLines() kKines {
	return kKines{
		RestAPIBase: api.NewRestAPIBase(),
	}
}

type kKines struct {
	api.RestAPIBase
}

func (k *kKines) KLines(symbol configuration.Symbol, intervalType configuration.IntervalType, date string) (*model.KLinesRes, error) {
	param := url.Values{
		"symbol":   {string(symbol)},
		"interval": {string(intervalType)},
		"date":     {date},
	}

	res, err := k.Get(param, "/v1/klines")
	if err != nil {
		return nil, err
	}

	kLinesRes := new(model.KLinesRes)
	err = json.Unmarshal(res, kLinesRes)
	if err != nil {
		return nil, fmt.Errorf("[KLines]error:%v,body:%s", err, res)
	}

	if kLinesRes.Status != 0 {
		return nil, fmt.Errorf("%v", kLinesRes.Messages)
	}

	return kLinesRes, nil
}
