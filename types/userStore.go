package types

import "context"

type UserStore interface {
	All(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, product User) error
	Update(ctx context.Context, id string, product User) (*User, error)
	Delete(ctx context.Context, id string) error
}
