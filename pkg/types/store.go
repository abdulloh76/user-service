package types

import "context"

type UserStore interface {
	GetUserDetails(ctx context.Context, id string) (*UserBody, error)
	UpdateUserCredentials(ctx context.Context, id string, credentials UpdateCredentialsDto) error
	UpdatePassword(ctx context.Context, id string, password string) error
	UpdateAddress(ctx context.Context, id string, user AddressModel) error
	DeleteUser(ctx context.Context, id string) error
}

type AuthStore interface {
	CreateUser(ctx context.Context, userCredentials SignInCredentials) (string, error)
	GetUser(ctx context.Context, email, passwordHash string) (*User, error)
}
