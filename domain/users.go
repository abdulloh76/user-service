package domain

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/abdulloh76/user-service/types"
)

var (
	ErrJsonUnmarshal = errors.New("failed to parse user from request body")
	ErrUserNotFound  = errors.New("user with given ID not found")
)

type Users struct {
	store types.UserStore
}

func NewUsersDomain(d types.UserStore) *Users {
	return &Users{
		store: d,
	}
}

func (u *Users) GetUser(ctx context.Context, id string) (*types.User, error) {
	user, err := u.store.Get(ctx, id)
	if user.ID == "" {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *Users) AllUsers(ctx context.Context) ([]types.User, error) {
	// todo add dto for getAll
	allUsers, err := u.store.All(ctx)
	if err != nil {
		return allUsers, err
	}

	return allUsers, nil
}

func (u *Users) CreateUser(ctx context.Context, body []byte) (*types.User, error) {
	// todo add dto for create
	user := types.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, ErrJsonUnmarshal
	}

	err := u.store.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *Users) ModifyUser(ctx context.Context, id string, body []byte) (*types.User, error) {
	modifiedUser := types.User{}
	if err := json.Unmarshal(body, &modifiedUser); err != nil {
		return nil, ErrJsonUnmarshal
	}

	updatedUser, err := u.store.Update(ctx, id, modifiedUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *Users) DeleteUser(ctx context.Context, id string) error {
	err := u.store.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
