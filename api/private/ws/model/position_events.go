package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
	"time"
)

func NewPositionEventsReq(command consts.WebSocketCommand, channel consts.WebSocketChannel) PositionEventsReq {
	return PositionEventsReq{
		WebsocketRequestCommon: model.WebsocketRequestCommon{
			Command: command,
			Channel: channel,
		},
	}
}

type PositionEventsReq struct {
	model.WebsocketRequestCommon
}

type PositionEventsRes struct {
	model.PrivateWebsocketResponseCommon
	PositionID   int64           `json:"positionId"`
	Symbol       consts.Symbol   `json:"symbol"`
	Side         consts.Side     `json:"side"`
	Size         decimal.Decimal `json:"size"`
	OrderdSize   decimal.Decimal `json:"orderdSize"`
	Price        decimal.Decimal `json:"price"`
	LossGain     decimal.Decimal `json:"lossGain"`
	Leverage     decimal.Decimal `json:"leverage"`
	LosscutPrice decimal.Decimal `json:"losscutPrice"`
	Timestamp    time.Time       `json:"timestamp"`
}
