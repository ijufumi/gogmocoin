package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoad_ReadsFromEnvironment verifies that load() reads the configuration
// from the process environment. load() also calls godotenv.Load(), which does
// not override variables already present in the environment, so the values set
// here take precedence.
func TestLoad_ReadsFromEnvironment(t *testing.T) {
	t.Setenv("DEBUG_MODE", "true")
	t.Setenv("API_KEY", "test-api-key")
	t.Setenv("API_SECRET", "test-api-secret")

	cfg := load()

	assert.True(t, cfg.DebugMode)
	assert.Equal(t, "test-api-key", cfg.APIKey)
	assert.Equal(t, "test-api-secret", cfg.APISecret)
}

// TestLoad_DebugModeDisabled confirms DEBUG_MODE=false is parsed as false.
func TestLoad_DebugModeDisabled(t *testing.T) {
	t.Setenv("DEBUG_MODE", "false")
	t.Setenv("API_KEY", "k")
	t.Setenv("API_SECRET", "s")

	cfg := load()

	assert.False(t, cfg.DebugMode)
}

// TestLoad_InvalidDebugModeFallsBackToDefault documents that a value that
// cannot be parsed makes load() return a zero-value config rather than panic.
func TestLoad_InvalidDebugModeFallsBackToDefault(t *testing.T) {
	t.Setenv("DEBUG_MODE", "not-a-bool")

	cfg := load()

	assert.False(t, cfg.DebugMode)
	assert.Empty(t, cfg.APIKey)
	assert.Empty(t, cfg.APISecret)
}
