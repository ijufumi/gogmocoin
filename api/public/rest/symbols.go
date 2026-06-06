package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/public/rest/model"
	"net/url"
)

// Symbols is the client interface for the trading rules endpoint.
type Symbols interface {
	Symbols() (*model.SymbolsRes, error)
	SymbolsWithContext(ctx context.Context) (*model.SymbolsRes, error)
}

func newSymbols() symbols {
	return symbols{
		RestAPIBase: api.NewRestAPIBase(),
	}
}

type symbols struct {
	api.RestAPIBase
}

// Symbols retrieves the trading rules for all symbols using a background context.
func (t *symbols) Symbols() (*model.SymbolsRes, error) {
	return t.SymbolsWithContext(context.Background())
}

// SymbolsWithContext retrieves the trading rules for all symbols using the provided context.
func (t *symbols) SymbolsWithContext(ctx context.Context) (*model.SymbolsRes, error) {
	res, err := t.Get(ctx, url.Values{}, "/v1/symbols")
	if err != nil {
		return nil, err
	}

	symbolsRes := new(model.SymbolsRes)
	err = json.Unmarshal(res, symbolsRes)
	if err != nil {
		return nil, fmt.Errorf("[Symbols]error:%v,body:%s", err, res)
	}
	if !symbolsRes.Success() {
		return nil, symbolsRes.Error()
	}

	return symbolsRes, nil
}
