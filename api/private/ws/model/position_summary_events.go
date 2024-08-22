package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
	"time"
)

type PositionSummaryEventsRes struct {
	model.PrivateWebsocketResponseCommon
	Symbol              consts.Symbol   `json:"symbol"`
	Side                consts.Side     `json:"side"`
	AveragePositionRate decimal.Decimal `json:"averagePositionRate"`
	PositionLossGain    decimal.Decimal `json:"positionLossGain"`
	SumOrderQuantity    decimal.Decimal `json:"sumOrderQuantity"`
	SumPositionQuantity decimal.Decimal `json:"sumPositionQuantity"`
	Timestamp           time.Time       `json:"timestamp"`
}
