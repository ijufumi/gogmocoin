package private

import (
	"api_client/api/private/internal/connect"
	"api_client/api/private/model"
	"encoding/json"
	"fmt"
)

// CancelOrder ...
type CancelOrder interface {
	CancelOrder(orderID int64) error
}

type cancelOrder struct {
	con *connect.Connection
}

// CancelOrder ...
func (c *cancelOrder) CancelOrder(orderID int64) error {
	req := model.CancelOrderReq{OrderID: orderID}
	res, err := c.con.Post(req, "/v1/cancelOrder")
	if err != nil {
		return err
	}

	cancelOrderRes := new(model.CancelOrderRes)
	err = json.Unmarshal(res, cancelOrderRes)
	if err != nil {
		return fmt.Errorf("[CancelOrder]error:%v,body:%s", err, res)
	}

	if len(cancelOrderRes.Messages) != 0 {
		return fmt.Errorf("%v", cancelOrderRes.Messages)
	}
	return nil

}
