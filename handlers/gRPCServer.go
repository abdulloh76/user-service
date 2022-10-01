package handlers

import (
	"context"

	"github.com/abdulloh76/user-service/handlers/userGrpc"
)

// GRPCServer ...
type GRPCServer struct {
	userGrpc.UnimplementedGetServer
}

func (s *GRPCServer) GetUserDetails(ctx context.Context, req *userGrpc.GetRequest) (*userGrpc.GetResponse, error) {
	return &userGrpc.GetResponse{
		Name:  "Sherlock",
		Email: "sher@lock.com",
	}, nil
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *userGrpc.GetRequest) (*userGrpc.GetResponse, error) {
	return &userGrpc.GetResponse{
		Name:  "Sherlock",
		Email: "sher@lock.com",
	}, nil
}
