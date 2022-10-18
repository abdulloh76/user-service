package domain

import (
	"context"
	"encoding/json"

	"github.com/abdulloh76/user-service/pkg/types"
	"github.com/abdulloh76/user-service/pkg/utils"
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
	user, err := u.store.GetUserDetails(ctx, id)
	if user.ID == "" {
		return nil, utils.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	// todo map the return object so password won't be sent with the user details
	return user, nil
}

func (u *Users) UpdateUserCredentials(ctx context.Context, id string, body []byte) error {
	credentials := types.UpdateCredentialsDto{}
	if err := json.Unmarshal(body, &credentials); err != nil {
		return utils.ErrJsonUnmarshal
	}

	err := u.store.UpdateUserCredentials(ctx, id, credentials)
	if err != nil {
		return err
	}

	return nil
}

func (u *Users) UpdatePassword(ctx context.Context, id string, body []byte) error {
	passwords := types.UpdatePasswordDto{}
	if err := json.Unmarshal(body, &passwords); err != nil {
		return utils.ErrJsonUnmarshal
	}

	// todo check that current password hash matches with the one from db

	err := u.store.UpdatePassword(ctx, id, passwords.NewPassword)
	if err != nil {
		return err
	}

	return nil
}

func (u *Users) UpdateAddress(ctx context.Context, id string, body []byte) error {
	newAddress := types.AddressModel{}
	if err := json.Unmarshal(body, &newAddress); err != nil {
		return utils.ErrJsonUnmarshal
	}

	err := u.store.UpdateAddress(ctx, id, newAddress)
	if err != nil {
		return err
	}

	return nil
}

func (u *Users) DeleteUser(ctx context.Context, id string) error {
	err := u.store.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
