package private

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/api/common/api"
	"net/url"

	"github.com/ijufumi/gogmocoin/api/private/model"
)

// AccountMargin ...
type AccountMargin interface {
	AccountMargin() (*model.AccountMarginRes, error)
}

func newAccountMargin(apiKey, secretKey string) accountMargin {
	return accountMargin{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type accountMargin struct {
	api.RestAPIBase
}

func (a *accountMargin) AccountMargin() (*model.AccountMarginRes, error) {
	res, err := a.Get(url.Values{}, "/v1/account/margin")
	if err != nil {
		return nil, err
	}
	accountMarginRes := new(model.AccountMarginRes)
	err = json.Unmarshal(res, accountMarginRes)
	if err != nil {
		return nil, fmt.Errorf("[AccountMargin]error:%v,body:%s", err, res)
	}

	if len(accountMarginRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", accountMarginRes.Messages)
	}

	return accountMarginRes, nil
}
