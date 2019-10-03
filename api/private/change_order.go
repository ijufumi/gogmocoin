package private

import (
	"encoding/json"
	"fmt"

	"github.com/ijufumi/gogmocoin/api/private/internal/connect"
	"github.com/ijufumi/gogmocoin/api/private/model"
	"github.com/shopspring/decimal"
)

// ChangeOrder ...
type ChangeOrder interface {
	ChangeOrder(orderID int64, price decimal.Decimal) (*model.ChangeOrderRes, error)
}

type changeOrder struct {
	con *connect.Connection
}

func (c *changeOrder) ChangeOrder(orderID int64, price decimal.Decimal) (*model.ChangeOrderRes, error) {
	request := model.ChangeOrderReq{
		OrderID: orderID,
		Price:   price,
	}

	res, err := c.con.Post(request, "/v1/changeOrder")
	if err != nil {
		return nil, err
	}

	response := new(model.ChangeOrderRes)
	err = json.Unmarshal(res, response)
	if err != nil {
		return nil, fmt.Errorf("[ChangeOrder]error:%v,body:%s", err, res)
	}

	if len(response.Messages) != 0 {
		return nil, fmt.Errorf("%v", response.Messages)
	}
	return response, nil
}
