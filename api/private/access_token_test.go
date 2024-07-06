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
	client := New()
	for _, method := range methods {
		// get access token, next extend access token, next delete access token
		tokenRes, err := client.AccessToken(method, gotToken)
		assert.NoError(t, err)

		if method == http.MethodPost {
			assert.NotEqual(t, tokenRes.Data, "")
			gotToken = tokenRes.Data
		}

		log.Printf("[%s] token: %s", method, tokenRes.Data)
	}
}
