package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/ijufumi/gogmocoin/v2/api/common/value"
	"github.com/shopspring/decimal"
)

type KLinesRes struct {
	Data []struct {
		OpenTime value.TimeInMillis `json:"openTime"`
		Open     decimal.Decimal    `json:"open"`
		High     decimal.Decimal    `json:"high"`
		Low      decimal.Decimal    `json:"low"`
		Close    decimal.Decimal    `json:"close"`
		Volume   decimal.Decimal    `json:"volume"`
	}
	model.ResponseCommon
}
