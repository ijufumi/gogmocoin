package private

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAccessToken(t *testing.T) {
	godotenv.Load("../../.env.sample")
	assert.NotEqual(t, os.Getenv("API_KEY"), "")

	// test all methods
	methods := []string{
		http.MethodPost,   // as get access token
		http.MethodPut,    // as extend access token
		http.MethodDelete, // as delete access token
	}

	gotToken := ""

	for _, method := range methods {
		client := New()
		token, err := client.AccessToken(method, gotToken)
		assert.NoError(t, err)
		if method == http.MethodPost {
			assert.NotEqual(t, token, "")
			gotToken = token
		} else {
			assert.Equal(t, token, gotToken)
		}

		log.Printf("[%s] token:%s", method, token)
	}
}
