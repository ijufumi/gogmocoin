package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/api/common/api"
	"github.com/ijufumi/gogmocoin/api/public/rest/model"
	"net/url"
)

// Symbols ...
type Symbols interface {
	Symbols() (*model.SymbolsRes, error)
}

func newSymbols() symbols {
	return symbols{
		RestAPIBase: api.NewRestAPIBase(),
	}
}

type symbols struct {
	api.RestAPIBase
}

func (t *symbols) Symbols() (*model.SymbolsRes, error) {
	res, err := t.Get(url.Values{}, "/v1/symbols")
	if err != nil {
		return nil, err
	}

	symbolsRes := new(model.SymbolsRes)
	err = json.Unmarshal(res, symbolsRes)
	if err != nil {
		return nil, fmt.Errorf("[Symbols]error:%v,body:%s", err, res)
	}
	if len(symbolsRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", symbolsRes.Messages)
	}

	return symbolsRes, nil
}
