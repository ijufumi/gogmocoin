package model

import (
	"encoding/json"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func Test_MarshalOrderBooksResponse_1(t *testing.T) {
	f, err := os.Open("orderbooks_test_1.json")
	assert.Nil(t, err)
	b, err := io.ReadAll(f)
	assert.Nil(t, err)

	var response OrderBooksRes
	err = json.Unmarshal(b, &response)
	assert.Nil(t, err)

	assert.Equal(t, consts.SymbolBTC, response.Data.Symbol)
}
