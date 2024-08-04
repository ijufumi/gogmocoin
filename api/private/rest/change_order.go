package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"github.com/shopspring/decimal"
)

// ChangeOrder ...
type ChangeOrder interface {
	ChangeOrder(orderID int64, price decimal.Decimal) (*model.ChangeOrderRes, error)
}

func newChangeOrder(apiKey, secretKey string) changeOrder {
	return changeOrder{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type changeOrder struct {
	api.RestAPIBase
}

func (c *changeOrder) ChangeOrder(orderID int64, price decimal.Decimal) (*model.ChangeOrderRes, error) {
	request := model.ChangeOrderReq{
		OrderID: orderID,
		Price:   price,
	}

	res, err := c.Post(request, "/v1/changeOrder")
	if err != nil {
		return nil, err
	}

	response := new(model.ChangeOrderRes)
	err = json.Unmarshal(res, response)
	if err != nil {
		return nil, fmt.Errorf("[ChangeOrder]error:%v,body:%s", err, res)
	}

	if !response.Success() {
		return nil, response.Error()
	}
	return response, nil
}
