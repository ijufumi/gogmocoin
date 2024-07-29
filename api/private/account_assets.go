package private

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/model"
	"net/url"
)

// AccountAssets ...
type AccountAssets interface {
	AccountAssets() (*model.AccountAssetsRes, error)
}

func newAccountAssets(apiKey, secretKey string) accountAssets {
	return accountAssets{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type accountAssets struct {
	api.RestAPIBase
}

func (a *accountAssets) AccountAssets() (*model.AccountAssetsRes, error) {
	res, err := a.Get(url.Values{}, "/v1/account/assets")
	if err != nil {
		return nil, err
	}

	accountAssetsRes := new(model.AccountAssetsRes)
	err = json.Unmarshal(res, accountAssetsRes)
	if err != nil {
		return nil, fmt.Errorf("[AccountAssets]error:%v,body:%s", err, res)
	}

	if len(accountAssetsRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", accountAssetsRes.Messages)
	}

	return accountAssetsRes, nil
}
