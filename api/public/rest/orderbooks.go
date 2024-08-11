package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/public/rest/model"
	"net/url"
)

// OrderBooks ...
type OrderBooks interface {
	OrderBooks(symbol consts.Symbol) (*model.OrderBooksRes, error)
	OrderBooksWithContext(ctx context.Context, symbol consts.Symbol) (*model.OrderBooksRes, error)
}

func newOrderBooks() orderBooks {
	return orderBooks{
		RestAPIBase: api.NewRestAPIBase(),
	}
}

type orderBooks struct {
	api.RestAPIBase
}

// OrderBooks ...
func (o *orderBooks) OrderBooks(symbol consts.Symbol) (*model.OrderBooksRes, error) {
	return o.OrderBooksWithContext(context.Background(), symbol)
}

// OrderBooksWithContext ...
func (o *orderBooks) OrderBooksWithContext(ctx context.Context, symbol consts.Symbol) (*model.OrderBooksRes, error) {
	param := url.Values{
		"symbol": {string(symbol)},
	}

	res, err := o.Get(ctx, param, "/v1/orderbooks")
	if err != nil {
		return nil, err
	}

	orderBooksRes := new(model.OrderBooksRes)
	err = json.Unmarshal(res, orderBooksRes)
	if err != nil {
		return nil, fmt.Errorf("[OrderBooks]error:%v,body:%s", err, res)
	}

	if !orderBooksRes.Success() {
		return nil, orderBooksRes.Error()
	}

	return orderBooksRes, nil

}
