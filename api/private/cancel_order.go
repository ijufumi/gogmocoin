package private

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/model"
)

// CancelOrder ...
type CancelOrder interface {
	CancelOrder(orderID int64) error
}

func newCancelOrder(apiKey, secretKey string) cancelOrder {
	return cancelOrder{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type cancelOrder struct {
	api.RestAPIBase
}

// CancelOrder ...
func (c *cancelOrder) CancelOrder(orderID int64) error {
	req := model.CancelOrderReq{OrderID: orderID}
	res, err := c.Post(req, "/v1/cancelOrder")
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
