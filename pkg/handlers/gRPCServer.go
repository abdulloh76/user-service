package handlers

import (
	"context"

	"github.com/abdulloh76/user-service/pkg/domain"
	"github.com/abdulloh76/user-service/pkg/handlers/userGrpc"
)

type GRPCServer struct {
	users *domain.Users
	userGrpc.UnimplementedUserServer
}

func NewGRPCServer(d *domain.Users) *GRPCServer {
	return &GRPCServer{
		users: d,
	}
}

func (s *GRPCServer) GetUserDetails(ctx context.Context, req *userGrpc.GetRequest) (*userGrpc.UserDetails, error) {
	id := req.GetId()

	user, err := s.users.GetUser(ctx, id)

	if err != nil {
		return nil, err
	}

	return &userGrpc.UserDetails{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Address: &userGrpc.Address{
			Street:   user.Address.Street,
			City:     user.Address.City,
			PostCode: user.Address.PostCode,
			Country:  user.Address.Country,
		},
	}, nil
}

func (s *GRPCServer) GetUserAddress(ctx context.Context, req *userGrpc.GetRequest) (*userGrpc.Address, error) {
	id := req.GetId()

	user, err := s.users.GetUser(ctx, id)

	if err != nil {
		return nil, err
	}

	return &userGrpc.Address{
		Street:   user.Address.Street,
		City:     user.Address.City,
		PostCode: user.Address.PostCode,
		Country:  user.Address.Country,
	}, nil
}
