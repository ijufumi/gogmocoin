package model

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/common/model"
	"github.com/shopspring/decimal"
)

type AccountTradingVolumeRes struct {
	model.ResponseCommon
	Data struct {
		JpyVolume decimal.Decimal  `json:"jpyVolume"`
		TierLevel consts.TierLevel `json:"tierLevel"`
		Limit     []struct {
			consts.Symbol      `json:"symbol"`
			TodayLimitOpenSize decimal.Decimal `json:"todayLimitOpenSize"`
			TakerFee           decimal.Decimal `json:"takerFee"`
			MakerFee           decimal.Decimal `json:"makerFee"`
		} `json:"limit"`
	}
}
