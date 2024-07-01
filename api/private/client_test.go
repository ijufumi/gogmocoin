package private

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var (
	SYMBOLS = []string{"ASTR", "BTC_JPY", "ETH_JPY", "BCH_JPY", "LTC_JPY", "XRP_JPY", "DOT_JPY", "ATOM_JPY", "ADA_JPY", "LINK_JPY", "DOGE_JPY", "SOL_JPY", "BTC", "ETH", "BCH", "LTC", "XRP", "XEM", "XLM", "BAT", "XTZ", "QTUM", "ENJ", "DOT", "ATOM", "MKR", "DAI", "XYM", "MONA", "FCR", "ADA", "LINK", "DOGE", "SOL"}
)

func TestParse(t *testing.T) {

	for _, v := range SYMBOLS {
		fmt.Printf("v: %+v\n", v)
	}
}

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

	for _, v := range SYMBOLS {
		res, err := client.LastExecutions(configuration.Symbol(v), page, count)
		assert.NoError(t, err)

		log.Printf("result: %+v\n", res)
	}
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
