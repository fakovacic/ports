package middleware

import (
	"context"

	"github.com/fakovacic/ports/internal/ports"
)

type loggingMiddleware struct {
	service string
	config  *ports.Config
	next    ports.Service
}

func NewLoggingMiddleware(config *ports.Config, next ports.Service) ports.Service {
	m := loggingMiddleware{
		service: ports.ServiceName,
		config:  config,
		next:    next,
	}

	return &m
}

func (m *loggingMiddleware) Create(ctx context.Context, input *ports.Port) (*ports.Port, error) {
	m.config.Log.Info().
		Str("service", m.service).
		Str("method", "Create").
		Interface("input", input).
		Msg("service request")

	model, err := m.next.Create(ctx, input)

	m.config.Log.Info().
		Str("service", m.service).
		Str("method", "Create").
		Interface("model", model).
		Err(err).
		Msg("service response")

	//nolint
	return model, err
}

func (m *loggingMiddleware) Update(ctx context.Context, id string, input *ports.Port) (*ports.Port, error) {
	m.config.Log.Info().
		Str("service", m.service).
		Str("method", "Update").
		Str("id", id).
		Interface("input", input).
		Msg("service request")

	model, err := m.next.Update(ctx, id, input)

	m.config.Log.Info().
		Str("service", m.service).
		Str("method", "Update").
		Interface("model", model).
		Err(err).
		Msg("service response")

	//nolint
	return model, err
}
