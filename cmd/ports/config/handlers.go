package config

import (
	"github.com/fakovacic/ports/internal/ports"
	"github.com/fakovacic/ports/internal/ports/handlers/http"
)

func NewHandlers(c *ports.Config, service ports.Service) http.Handler {
	return *http.New(c, service)
}
