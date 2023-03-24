package ports

import "context"

//go:generate moq -out ./mocks/store.go -pkg mocks  . Store
type Store interface {
	Create(context.Context, *Port) error
	Update(context.Context, string, *Port) error
}
