package model

import (
	"encoding/json"
	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func Test_MarshalAccountTradingVolumeResponse_Success(t *testing.T) {
	f, err := os.Open("account_trading_volume_test_success.json")
	assert.Nil(t, err)
	b, err := io.ReadAll(f)
	assert.Nil(t, err)

	var response AccountTradingVolumeRes
	err = json.Unmarshal(b, &response)
	assert.Nil(t, err)

	assert.Equal(t, configuration.SymbolBTCJPY, response.Data.Limit[0].Symbol)
	assert.Equal(t, configuration.SymbolBTC, response.Data.Limit[1].Symbol)
}

func Test_MarshalAccountTradingVolumeResponse_Failure(t *testing.T) {
	f, err := os.Open("account_trading_volume_test_failure.json")
	assert.Nil(t, err)
	b, err := io.ReadAll(f)
	assert.Nil(t, err)

	var response AccountTradingVolumeRes
	err = json.Unmarshal(b, &response)
	assert.Equal(t, configuration.ErrUnsupportedTierLevel, err)
}
