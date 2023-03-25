package ports

import (
	"context"

	"github.com/fakovacic/ports/internal/errors"
)

func (s *service) Create(ctx context.Context, m *Port) (*Port, error) {
	err := validateCreate(ctx, m)
	if err != nil {
		return nil, err
	}

	err = s.store.Create(ctx, m)
	if err != nil {
		return nil, errors.Wrapf(err, "create port")
	}

	return m, nil
}

func validateCreate(_ context.Context, m *Port) error {
	if m.Name == "" {
		return errors.Wrap("name is required")
	}

	// add additional validation

	return nil
}
