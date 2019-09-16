package model

import (
	"api_client/api/common/model"

	"github.com/shopspring/decimal"
)

// ChangeLosscutPriceReq ...
type ChangeLosscutPriceReq struct {
	PositionID   int64           `json:"positionId"`
	LosscutPrice decimal.Decimal `json:"losscutPrice"`
}

// ChangeLosscutPriceRes ...
type ChangeLosscutPriceRes struct {
	model.ResponseCommon
}
