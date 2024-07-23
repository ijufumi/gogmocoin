package model

import (
	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

type AccountTradingVolumeRes struct {
	model.ResponseCommon
	Data struct {
		JpyVolume decimal.Decimal         `json:"jpyVolume"`
		TierLevel configuration.TierLevel `json:"tierLevel"`
		Limit     []struct {
			configuration.Symbol `json:"symbol"`
			TodayLimitOpenSize   decimal.Decimal `json:"todayLimitOpenSize"`
			TakerFee             decimal.Decimal `json:"takerFee"`
			MakerFee             decimal.Decimal `json:"makerFee"`
		} `json:"limit"`
	}
}
