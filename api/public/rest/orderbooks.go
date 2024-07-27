package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/api/common/api"
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"net/url"

	"github.com/ijufumi/gogmocoin/api/public/rest/model"
)

// OrderBooks ...
type OrderBooks interface {
	OrderBooks(symbol consts.Symbol) (*model.OrderBooksRes, error)
}

func newOrderBooks() orderBooks {
	return orderBooks{
		RestAPIBase: api.NewRestAPIBase(),
	}
}

type orderBooks struct {
	api.RestAPIBase
}

func (o *orderBooks) OrderBooks(symbol consts.Symbol) (*model.OrderBooksRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	res, err := o.Get(param, "/v1/orderbooks")
	if err != nil {
		return nil, err
	}

	orderBooksRes := new(model.OrderBooksRes)
	err = json.Unmarshal(res, orderBooksRes)
	if err != nil {
		return nil, fmt.Errorf("[OrderBooks]error:%v,body:%s", err, res)
	}

	if len(orderBooksRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", orderBooksRes.Messages)
	}

	return orderBooksRes, nil

}
