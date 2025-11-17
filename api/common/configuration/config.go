package configuration

import (
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	DebugMode bool   `env:"DEBUG_MODE"`
	APIKey    string `env:"API_KEY"`
	APISecret string `env:"API_SECRET"`
}

func load() *config {
	c, err := env.ParseAs[config]()
	if err != nil {
		// Return default config if parsing fails
		return &config{}
	}
	return &c
}

var c = load()

// IsDebug ...
func IsDebug() bool {
	return c.DebugMode
}

func APIKey() string {
	return c.APIKey
}

func APISecret() string {
	return c.APISecret
}
