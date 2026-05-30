package configuration

import (
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type config struct {
	DebugMode bool   `env:"DEBUG_MODE"`
	APIKey    string `env:"API_KEY"`
	APISecret string `env:"API_SECRET"`
}

func load() *config {
	// Load .env for convenience, but godotenv.Load does not override variables
	// that are already present in the process environment, so a caller that sets
	// its own environment (Vault, Kubernetes secrets, os.Setenv, ...) keeps
	// precedence.
	_ = godotenv.Load()
	c, err := env.ParseAs[config]()
	if err != nil {
		// Return default config if parsing fails
		return &config{}
	}
	return &c
}

var (
	once   sync.Once
	cached *config
)

// get loads the configuration lazily on first access. This avoids reading the
// environment (and the .env file) as an import-time side effect, which a
// library must not impose on its callers.
func get() *config {
	once.Do(func() {
		cached = load()
	})
	return cached
}

// IsDebug reports whether debug output is enabled (DEBUG_MODE).
func IsDebug() bool {
	return get().DebugMode
}

// APIKey returns the API key read from the environment (API_KEY).
func APIKey() string {
	return get().APIKey
}

// APISecret returns the API secret read from the environment (API_SECRET).
func APISecret() string {
	return get().APISecret
}
