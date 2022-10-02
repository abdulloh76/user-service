package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/abdulloh76/user-service/domain"
	"github.com/abdulloh76/user-service/handlers"
	"github.com/abdulloh76/user-service/handlers/userGrpc"
	"github.com/abdulloh76/user-service/store"
	"google.golang.org/grpc"
)

func main() {
	PORT := os.Getenv("PORT")
	DYNAMODB_PORT := os.Getenv("DYNAMODB_PORT")
	tableName := "users"

	dynamoDB := store.NewDynamoDBStore(context.TODO(), DYNAMODB_PORT, tableName)
	domain := domain.NewUsersDomain(dynamoDB)
	grpcServer := handlers.NewGRPCServer(domain)

	s := grpc.NewServer()
	userGrpc.RegisterUserServer(s, grpcServer)

	l, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
