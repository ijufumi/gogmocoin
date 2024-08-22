package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
)

// AccountMargin ...
type AccountMargin interface {
	AccountMargin() (*model.AccountMarginRes, error)
	AccountMarginWithContext(ctx context.Context) (*model.AccountMarginRes, error)
}

func newAccountMargin(apiKey, secretKey string) accountMargin {
	return accountMargin{
		apiBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type accountMargin struct {
	apiBase api.RestAPIBase
}

// AccountMargin ...
func (a *accountMargin) AccountMargin() (*model.AccountMarginRes, error) {
	return a.AccountMarginWithContext(context.Background())
}

// AccountMarginWithContext ...
func (a *accountMargin) AccountMarginWithContext(ctx context.Context) (*model.AccountMarginRes, error) {
	res, err := a.apiBase.Get(ctx, url.Values{}, "/v1/account/margin")
	if err != nil {
		return nil, err
	}
	accountMarginRes := new(model.AccountMarginRes)
	err = json.Unmarshal(res, accountMarginRes)
	if err != nil {
		return nil, fmt.Errorf("[AccountMargin]error:%v,body:%s", err, res)
	}

	if !accountMarginRes.Success() {
		return nil, accountMarginRes.Error()
	}

	return accountMarginRes, nil
}
