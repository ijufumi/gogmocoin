package private

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ijufumi/gogmocoin/api/private/internal/connect"
	"github.com/ijufumi/gogmocoin/api/private/model"
)

// AccountMargin ...
type AccountMargin interface {
	AccountMargin() (*model.AccountMarginRes, error)
}

type accountMargin struct {
	con *connect.Connection
}

func (a *accountMargin) AccountMargin() (*model.AccountMarginRes, error) {
	res, err := a.con.Get(url.Values{}, "/v1/account/margin")
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
