package handlers

import (
	"context"

	"github.com/abdulloh76/user-service/domain"
	"github.com/abdulloh76/user-service/handlers/userGrpc"
)

// GRPCServer ...
type GRPCServer struct {
	users *domain.Users
	userGrpc.UnimplementedUserServer
}

func NewGRPCServer(d *domain.Users) *GRPCServer {
	return &GRPCServer{
		users: d,
	}
}

func (s *GRPCServer) GetUserDetails(ctx context.Context, req *userGrpc.GetRequest) (*userGrpc.GetResponse, error) {
	return &userGrpc.GetResponse{
		Name:  "Sherlock",
		Email: "sher@lock.com",
		Address: &userGrpc.Address{
			Street:   "221B Baker Street",
			City:     "London",
			PostCode: "NR24 5WQ",
			Country:  "UK",
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
