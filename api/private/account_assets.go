package private

import (
	"api_client/api/private/internal/connect"
	"api_client/api/private/model"
	"encoding/json"
	"fmt"
	"net/url"
)

// AccountAssets ...
type AccountAssets interface {
	AccountAssets() (*model.AccountAssetsRes, error)
}

type accountAssets struct {
	con *connect.Connection
}

func (a *accountAssets) AccountAssets() (*model.AccountAssetsRes, error) {
	res, err := a.con.Get(url.Values{}, "/v1/account/assets")
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
