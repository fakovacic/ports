package config

import (
	"github.com/fakovacic/ports/internal/ports"
	"github.com/fakovacic/ports/internal/ports/store/inmem"
)

func NewStore() (ports.Store, error) {
	return inmem.New(), nil
}
