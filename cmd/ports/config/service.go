package config

import (
	"github.com/fakovacic/ports/internal/ports"
	"github.com/fakovacic/ports/internal/ports/middleware"
)

func NewService(c *ports.Config, store ports.Store) (ports.Service, error) {
	service := ports.New(c, store)
	service = middleware.NewLoggingMiddleware(c, service)

	return service, nil
}
