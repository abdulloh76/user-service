package utils

import "errors"

var (
	ErrEmptyBody     = errors.New("empty request body")
	ErrJsonUnmarshal = errors.New("failed to parse user from request body")
	ErrUserNotFound  = errors.New("user with given ID not found")

	ErrUserNotExists = errors.New("no user with this email, please create one")
	ErrWrongPassword = errors.New("password did not match")
	ErrWithDB        = errors.New("db error")

	ErrInvalidJWTMethod   = errors.New("invalid signing method")
	ErrInvalidTokenClaims = errors.New("token claims are not of type *tokenClaims")
)
