package private

import (
	"api_client/api/common/configuration"
	"api_client/api/private/internal/connect"
	"api_client/api/private/model"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type LastExecutions interface {
	LastExecutions(symbol configuration.Symbol, page, count int) (*model.LastExecutionsRes, error)
}

type lastExecutions struct {
	con *connect.Connection
}

func (l *lastExecutions) LastExecutions(symbol configuration.Symbol, page, count int) (*model.LastExecutionsRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	if page > 0 {
		param.Set("page", strconv.Itoa(page))
	}

	if count > 0 {
		param.Set("count", strconv.Itoa(count))
	}

	res, err := l.con.Get(param, "/v1/latestExecutions")
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
