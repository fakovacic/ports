package inmem

import (
	"context"

	"github.com/fakovacic/ports/internal/ports"
)

func (s *store) Create(ctx context.Context, model *ports.Port) error {
	s.data.Store(model.ID, model)

	return nil
}
