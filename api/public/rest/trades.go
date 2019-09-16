package rest

import (
	"api_client/api/common/configuration"
	"api_client/api/public/rest/internal/connect"
	"api_client/api/public/rest/model"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Trades interface {
	Trades(symbol configuration.Symbol, page, count int64) (*model.TradesRes, error)
}

type trades struct {
	con connect.Connection
}

func (t *trades) Trades(symbol configuration.Symbol, page, count int64) (*model.TradesRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	if page > 0 {
		param.Set("page", strconv.FormatInt(page, 10))
	}

	if count > 0 {
		param.Set("count", strconv.FormatInt(count, 10))
	}

	res, err := t.con.Get(param, "/v1/trades")
	if err != nil {
		return nil, err
	}

	tradesRes := new(model.TradesRes)
	err = json.Unmarshal(res, tradesRes)
	if err != nil {
		return nil, fmt.Errorf("[Trades]error:%v,body:%s", err, res)
	}
	if len(tradesRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", tradesRes.Messages)
	}

	return tradesRes, nil
}
