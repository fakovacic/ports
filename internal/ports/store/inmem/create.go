package inmem

import (
	"context"

	"github.com/fakovacic/ports/internal/ports"
)

func (s *store) Create(_ context.Context, model *ports.Port) error {
	s.data.Store(model.ID, model)

	return nil
}
