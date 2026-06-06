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

// OrderBooks is the client interface for the public order book endpoint.
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

// OrderBooks retrieves the current order book for the given symbol using a background context.
func (o *orderBooks) OrderBooks(symbol consts.Symbol) (*model.OrderBooksRes, error) {
	return o.OrderBooksWithContext(context.Background(), symbol)
}

// OrderBooksWithContext retrieves the current order book for the given symbol using the provided context.
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
