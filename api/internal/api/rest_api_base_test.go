package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestMaskSensitiveHeader verifies that credential headers are redacted in the
// returned copy while non-sensitive headers are preserved and the original
// header is left untouched.
func TestMaskSensitiveHeader(t *testing.T) {
	original := http.Header{}
	original.Set("API-KEY", "my-api-key")
	original.Set("API-SIGN", "my-signature")
	original.Set("API-TIMESTAMP", "1700000000000")
	original.Set("Content-Type", "application/json")

	masked := maskSensitiveHeader(original)

	assert.Equal(t, "[REDACTED]", masked.Get("API-KEY"))
	assert.Equal(t, "[REDACTED]", masked.Get("API-SIGN"))
	assert.Equal(t, "1700000000000", masked.Get("API-TIMESTAMP"), "non-credential headers must be preserved")
	assert.Equal(t, "application/json", masked.Get("Content-Type"))

	// The original header must not be mutated.
	assert.Equal(t, "my-api-key", original.Get("API-KEY"))
	assert.Equal(t, "my-signature", original.Get("API-SIGN"))
}

// TestMaskSensitiveHeader_NoCredentials ensures headers without credentials are
// returned unchanged and that empty/nil inputs do not panic.
func TestMaskSensitiveHeader_NoCredentials(t *testing.T) {
	h := http.Header{}
	h.Set("Content-Type", "application/json")
	masked := maskSensitiveHeader(h)
	assert.Equal(t, "application/json", masked.Get("Content-Type"))
	assert.Empty(t, masked.Get("API-KEY"))

	assert.NotPanics(t, func() {
		_ = maskSensitiveHeader(http.Header{})
		_ = maskSensitiveHeader(nil)
	})
}

// TestMakeSign locks the signing input order (timestamp + method + path + body)
// so a refactor of makeSign cannot silently change the produced signature.
func TestMakeSign(t *testing.T) {
	secret := "test-secret"
	ts := "1700000000000"
	method := httpMethodPOST
	path := "/v1/order"
	body := `{"symbol":"BTC_JPY"}`

	got := makeSign(secret, ts, method, path, body)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(ts + string(method) + path + body))
	want := hex.EncodeToString(mac.Sum(nil))

	assert.Equal(t, want, got)
}

// TestMakeAuthHeader verifies the auth headers are populated, including the
// millisecond timestamp and a non-empty signature.
func TestMakeAuthHeader(t *testing.T) {
	r, err := http.NewRequest(http.MethodPost, "https://example.com/v1/order", nil)
	assert.NoError(t, err)

	makeAuthHeader("my-key", "my-secret", time.UnixMilli(1700000000000), r, httpMethodPOST, "/v1/order", "{}")

	assert.Equal(t, "1700000000000", r.Header.Get("API-TIMESTAMP"))
	assert.Equal(t, "my-key", r.Header.Get("API-KEY"))
	assert.NotEmpty(t, r.Header.Get("API-SIGN"))
}
