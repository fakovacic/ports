package ports

import (
	"context"

	"github.com/fakovacic/ports/internal/ports/errors"
)

func (s *service) Update(ctx context.Context, id string, m *Port) (*Port, error) {
	if id == "" {
		return nil, errors.BadRequest("id is empty")
	}

	err := validateUpdate(ctx, m)
	if err != nil {
		return nil, err
	}

	// can add check to get Port from store first
	// in order to verify if exist

	err = s.store.Update(ctx, id, m)
	if err != nil {
		return nil, errors.Wrapf(err, "update port")
	}

	return m, nil
}

func validateUpdate(_ context.Context, m *Port) error {
	if m.Name == "" {
		return errors.Wrap("name is required")
	}

	// add additional validation

	return nil
}
