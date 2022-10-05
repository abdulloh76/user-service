package types

import "context"

type UserStore interface {
	All(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user UserBody) error
	Update(ctx context.Context, id string, user UserBody) (*User, error)
	Delete(ctx context.Context, id string) error
}

type AuthStore interface {
	CreateUser(ctx context.Context, userCredentials UserCredentials) (string, error)
	GetUser(ctx context.Context, email, passwordHash string) (*User, error)
}
