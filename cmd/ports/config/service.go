package config

import (
	"github.com/fakovacic/ports/internal/ports"
)

func NewService(c *ports.Config, store ports.Store) (ports.Service, error) {
	return ports.New(c, store), nil
}
