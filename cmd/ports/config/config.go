package config

import (
	"os"

	"github.com/fakovacic/ports/internal/ports"
)

func NewConfig() (*ports.Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	return ports.NewConfig(env), nil
}
