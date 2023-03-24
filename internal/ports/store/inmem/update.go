package inmem

import (
	"context"

	"github.com/fakovacic/ports/internal/ports"
)

func (s *store) Update(ctx context.Context, id string, model *ports.Port) error {
	s.data.Store(id, model)

	return nil
}
