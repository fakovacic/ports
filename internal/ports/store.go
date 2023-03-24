package ports

import "context"

type Store interface {
	Create(context.Context, *Port) error
	Update(context.Context, string, *Port) error
}
