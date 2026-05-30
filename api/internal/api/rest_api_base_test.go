package api

import (
	"net/http"
	"testing"

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
