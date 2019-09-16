package rest

import (
	"api_client/api/common/configuration"
	"api_client/api/public/rest/internal/connect"
	"api_client/api/public/rest/model"
	"encoding/json"
	"fmt"
	"net/url"
)

type OrderBooks interface {
	OrderBooks(symbol configuration.Symbol) (*model.OrderBooksRes, error)
}

type orderbooks struct {
	con connect.Connection
}

func (o *orderbooks) OrderBooks(symbol configuration.Symbol) (*model.OrderBooksRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	res, err := o.con.Get(param, "/v1/orderbooks")
	if err != nil {
		return nil, err
	}

	orderbooksRes := new(model.OrderBooksRes)
	err = json.Unmarshal(res, orderbooksRes)
	if err != nil {
		return nil, fmt.Errorf("[OrderBooks]error:%v,body:%s", err, res)
	}

	if len(orderbooksRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", orderbooksRes.Messages)
	}

	return orderbooksRes, nil

}
