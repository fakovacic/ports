package parser

import (
	"context"
	"io"

	"github.com/fakovacic/ports/internal/ports"
)

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service
type Service interface {
	Parse(context.Context, io.Reader) error
}

func New(portsService ports.Service) Service {
	return &service{
		service: portsService,
	}
}

type service struct {
	service ports.Service
}
