package private

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountAssets(t *testing.T) {
	godotenv.Load("../../.env.sample")
	assert.NotEqual(t, os.Getenv("API_KEY"), "")

	client := New()
	res, err := client.AccountAssets()
	assert.NoError(t, err)

	log.Printf("result: %+v", res)
}

func TestGetActiveOrders(t *testing.T) {
	godotenv.Load("../../.env.sample")
	assert.NotEqual(t, os.Getenv("API_KEY"), "")

	no := 10

	client := New()
	res, err := client.ActiveOrders(configuration.SymbolBTCJPY, no)
	assert.NoError(t, err)

	log.Printf("result: %+v", res)
}

func TestGetLastExecutions(t *testing.T) {
	godotenv.Load("../../.env.sample")
	assert.NotEqual(t, os.Getenv("API_KEY"), "")

	page := 1
	count := 10

	client := New()
	res, err := client.LastExecutions(configuration.SymbolBTCJPY, page, count)
	assert.NoError(t, err)

	log.Printf("result: %+v", res)
}

func TestOrderAndCancel(t *testing.T) {
	godotenv.Load("../../.env.sample")
	assert.NotEqual(t, os.Getenv("API_KEY"), "")

	client := New()
	res, err := client.Order(
		configuration.SymbolBTCJPY,
		configuration.SideBUY,
		configuration.ExecutionTypeLIMIT,
		decimal.NewFromInt(10_000_000),
		decimal.NewFromFloat(0.001))
	assert.NoError(t, err)

	log.Printf("order result: %+v", res)
	orderId := res.Data

	time.Sleep(5 * time.Second)

	assert.NoError(t, client.CancelOrder(orderId))
}
