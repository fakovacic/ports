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

// NewLoggingMiddleware creates a new logging middleware.
func NewLoggingMiddleware(config *ports.Config, next ports.Service) ports.Service {
	m := loggingMiddleware{
		service: ports.ServiceName,
		config:  config,
		next:    next,
	}

	return &m
}

func (m *loggingMiddleware) Create(ctx context.Context, input *ports.Port) (*ports.Port, error) {
	reqID := ports.GetCtxStringVal(ctx, ports.ContextKeyRequestID)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Create").
		Interface("input", input).
		Msg("service request")

	model, err := m.next.Create(ctx, input)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Create").
		Interface("model", model).
		Err(err).
		Msg("service response")

	return model, err
}

func (m *loggingMiddleware) Update(ctx context.Context, id string, input *ports.Port) (*ports.Port, error) {
	reqID := ports.GetCtxStringVal(ctx, ports.ContextKeyRequestID)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Update").
		Str("id", id).
		Interface("input", input).
		Msg("service request")

	model, err := m.next.Update(ctx, id, input)

	m.config.Log.Info().
		Str("reqID", reqID).
		Str("service", m.service).
		Str("method", "Update").
		Interface("model", model).
		Err(err).
		Msg("service response")

	return model, err
}
