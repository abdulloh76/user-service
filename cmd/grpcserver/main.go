package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/abdulloh76/user-service/pkg/domain"
	"github.com/abdulloh76/user-service/pkg/handlers"
	"github.com/abdulloh76/user-service/pkg/handlers/userGrpc"
	"github.com/abdulloh76/user-service/pkg/store"
	"google.golang.org/grpc"
)

func main() {
	PORT := os.Getenv("PORT")
	tableName := "users"

	dynamoDB := store.NewDynamoDBStore(context.TODO(), tableName)
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
