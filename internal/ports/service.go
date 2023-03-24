package ports

import "context"

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service
type Service interface {
	Create(context.Context, *Port) (*Port, error)
	Update(context.Context, string, *Port) (*Port, error)
}

func New(c *Config, store Store) Service {
	return &service{
		config: c,
		store:  store,
	}
}

type service struct {
	config *Config
	store  Store
}
