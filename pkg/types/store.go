package types

import "context"

type UserStore interface {
	AllUsers(ctx context.Context) ([]User, error)
	GetUserDetails(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user UserBody) error
	UpdateUserDetails(ctx context.Context, id string, user UserBody) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

type AuthStore interface {
	CreateUser(ctx context.Context, userCredentials UserCredentials) (string, error)
	GetUser(ctx context.Context, email, passwordHash string) (*User, error)
}
