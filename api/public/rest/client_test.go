package rest

import (
	"log"
	"testing"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/stretchr/testify/assert"
)

func TestGetTicker(t *testing.T) {
	client := New()

	res, err := client.Ticker(configuration.SymbolNONE)
	assert.NoError(t, err)

	log.Printf("result: %+v", res)
}
