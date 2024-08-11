package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
)

// AccountAssets ...
type AccountAssets interface {
	AccountAssets() (*model.AccountAssetsRes, error)
	AccountAssetsWithContext(ctx context.Context) (*model.AccountAssetsRes, error)
}

func newAccountAssets(apiKey, secretKey string) accountAssets {
	return accountAssets{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type accountAssets struct {
	api.RestAPIBase
}

// AccountAssets ...
func (a *accountAssets) AccountAssets() (*model.AccountAssetsRes, error) {
	return a.AccountAssetsWithContext(context.Background())
}

// AccountAssetsWithContext ...
func (a *accountAssets) AccountAssetsWithContext(ctx context.Context) (*model.AccountAssetsRes, error) {
	res, err := a.Get(ctx, url.Values{}, "/v1/account/assets")
	if err != nil {
		return nil, err
	}

	accountAssetsRes := new(model.AccountAssetsRes)
	err = json.Unmarshal(res, accountAssetsRes)
	if err != nil {
		return nil, fmt.Errorf("[AccountAssets]error:%v,body:%s", err, res)
	}

	if !accountAssetsRes.Success() {
		return nil, accountAssetsRes.Error()
	}

	return accountAssetsRes, nil
}
