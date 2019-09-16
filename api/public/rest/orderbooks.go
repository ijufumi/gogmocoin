package rest

import (
	"encoding/json"
	"fmt"
	"gogmocoin/api/common/configuration"
	"gogmocoin/api/public/rest/internal/connect"
	"gogmocoin/api/public/rest/model"
	"net/url"
)

// OrderBooks ...
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
