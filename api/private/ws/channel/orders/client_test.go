package orders_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/ijufumi/gogmocoin/api/private/ws/channel/orders"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestClient_Connection(t *testing.T) {
	godotenv.Load("../../../../../.env.sample")
	assert.NotEqual(t, os.Getenv("API_KEY"), "")

	client := orders.New()
	assert.NotNil(t, client, "Client should not be nil")

	err := client.Subscribe()
	assert.NoError(t, err, "Subscription should not produce an error")

	// Simulate data reception and other assertions here
EXIT:
	for {
		select {
		case v := <-client.Receive():
			log.Printf("msg:%+v", v)

		case <-time.After(60 * time.Second):
			log.Println("timeout...")
			break EXIT
		}
	}

	err = client.Unsubscribe()
	assert.NoError(t, err, "Unsubscription should not produce an error")
}
