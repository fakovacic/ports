package inmem

import (
	"sync"

	"github.com/fakovacic/ports/internal/ports"
)

func New() ports.Store {
	s := &store{
		data: sync.Map{},
	}

	return s
}

type store struct {
	data sync.Map
}
