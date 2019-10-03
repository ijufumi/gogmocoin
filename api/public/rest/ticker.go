package rest

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/public/rest/internal/connect"
	"github.com/ijufumi/gogmocoin/api/public/rest/model"
)

// Ticker ...
type Ticker interface {
	Ticker(symbol configuration.Symbol) (*model.TickerRes, error)
}

type ticker struct {
	con connect.Connection
}

// Ticker ...
func (t ticker) Ticker(symbol configuration.Symbol) (*model.TickerRes, error) {
	param := url.Values{}

	if symbol != configuration.SymbolNONE {
		param.Set("symbol", string(symbol))
	}

	res, err := t.con.Get(param, "/v1/ticker")
	if err != nil {
		return nil, err
	}

	tickerRes := new(model.TickerRes)
	err = json.Unmarshal(res, tickerRes)

	if err != nil {
		return nil, fmt.Errorf("[Ticker]error:%v,body:%s", err, res)
	}

	if len(tickerRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", tickerRes.Messages)
	}

	return tickerRes, nil
}
