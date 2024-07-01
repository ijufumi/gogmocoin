package positions_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/ijufumi/gogmocoin/api/private/ws/channel/positions"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestClient_Connection(t *testing.T) {
	godotenv.Load("../../../../../.env.sample")
	assert.NotEqual(t, os.Getenv("API_KEY"), "")

	client := positions.New()
	assert.NotNil(t, client, "Client should not be nil")

	rxChan := client.Receive()
	assert.NotNil(t, rxChan, "Receive channel should not be nil")

	// Simulate data reception and other assertions here
	endTicker := time.NewTimer(60 * time.Second)
EXIT:
	select {
	case res := <-rxChan:
		assert.NotNil(t, res, "Received data should not be nil")
		// データ確認に関してはここに詳細なテストとアサートを追加してください。通常はJSONの構造や具体的なフィールドを確認します。
		log.Printf("received:%v", res)

	case <-endTicker.C:
		log.Println("timeout")
		break EXIT
	}

	err := client.Unsubscribe()
	assert.NoError(t, err, "Unsubscription should not produce an error")
}
