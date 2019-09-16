package private

import (
	"api_client/api/private/internal/connect"
	"api_client/api/private/model"
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

// ChangeLosscutPrice ...
type ChangeLosscutPrice interface {
	ChangeLosscutPrice(positionID int64, losscutPrice decimal.Decimal) (*model.ChangeLosscutPriceRes, error)
}

type changeLosscutPrice struct {
	con *connect.Connection
}

func (c *changeLosscutPrice) ChangeLosscutPrice(positionID int64, losscutPrice decimal.Decimal) (*model.ChangeLosscutPriceRes, error) {
	request := model.ChangeLosscutPriceReq{
		PositionID:   positionID,
		LosscutPrice: losscutPrice,
	}

	res, err := c.con.Post(request, "/v1/changeLosscutPrice")
	if err != nil {
		return nil, err
	}

	response := new(model.ChangeLosscutPriceRes)
	err = json.Unmarshal(res, response)
	if err != nil {
		return nil, fmt.Errorf("[ChangeLosscutPrice]error:%v,body:%s", err, res)
	}

	if len(response.Messages) != 0 {
		return nil, fmt.Errorf("%v", response.Messages)
	}
	return response, nil
}
