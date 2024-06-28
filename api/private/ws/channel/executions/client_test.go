package executions_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/ijufumi/gogmocoin/api/private/ws/channel/executions"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestClient_Connection(t *testing.T) {
	godotenv.Load("../../../../../.env.sample")
	assert.NotEqual(t, os.Getenv("API_KEY"), "")

	client := executions.New()
	assert.NotNil(t, client, "Client should not be nil")

	err := client.Subscribe()
	assert.NoError(t, err, "Subscription should not produce an error")

	rxChan := client.Receive()
	assert.NotNil(t, rxChan, "Receive channel should not be nil")

	// Simulate data reception and other assertions here
EXIT:
	select {
	case res := <-rxChan:
		assert.NotNil(t, res, "Received data should not be nil")
		// データ確認に関してはここに詳細なテストとアサートを追加してください。通常はJSONの構造や具体的なフィールドを確認します。
		log.Printf("received:%v", res)

	case <-time.After(60 * time.Second):
		log.Println("timeout")
		break EXIT
	}

	err = client.Unsubscribe()
	assert.NoError(t, err, "Unsubscription should not produce an error")
}
