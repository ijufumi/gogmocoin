package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
)

// ChangeLosscutPriceReq is the request of the change loss-cut price endpoint.
type ChangeLosscutPriceReq struct {
	PositionID   int64           `json:"positionId"`
	LosscutPrice decimal.Decimal `json:"losscutPrice"`
}

// ChangeLosscutPriceRes is the response of the change loss-cut price endpoint.
type ChangeLosscutPriceRes struct {
	model.ResponseCommon
}
