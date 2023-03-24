package ports

import (
	"os"

	"github.com/rs/zerolog"
)

const (
	ServiceName string = "ports"
)

func NewConfig(env string) *Config {
	return &Config{
		Env: env,
		Log: newLogger(),
	}
}

type Config struct {
	Log zerolog.Logger
	Env string
}

func newLogger() zerolog.Logger {
	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	return zerolog.New(os.Stdout).
		With().Timestamp().Logger()
}
