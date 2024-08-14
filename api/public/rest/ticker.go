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

// Ticker ...
type Ticker interface {
	Ticker(symbol consts.Symbol) (*model.TickerRes, error)
	TickerWithContext(ctx context.Context, symbol consts.Symbol) (*model.TickerRes, error)
}

func newTicker() ticker {
	return ticker{
		apiBase: api.NewRestAPIBase(),
	}
}

type ticker struct {
	apiBase api.RestAPIBase
}

// Ticker ...
func (t ticker) Ticker(symbol consts.Symbol) (*model.TickerRes, error) {
	return t.TickerWithContext(context.Background(), symbol)
}

// TickerWithContext ...
func (t ticker) TickerWithContext(ctx context.Context, symbol consts.Symbol) (*model.TickerRes, error) {
	param := url.Values{}

	if symbol != consts.SymbolNONE {
		param.Set("symbol", string(symbol))
	}

	res, err := t.apiBase.Get(ctx, param, "/v1/ticker")
	if err != nil {
		return nil, err
	}

	tickerRes := new(model.TickerRes)
	err = json.Unmarshal(res, tickerRes)

	if err != nil {
		return nil, fmt.Errorf("[Ticker]error:%v,body:%s", err, res)
	}

	if !tickerRes.Success() {
		return nil, tickerRes.Error()
	}

	return tickerRes, nil
}
