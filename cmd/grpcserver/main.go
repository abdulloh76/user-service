package main

import (
	"context"
	"log"
	"net"

	"github.com/abdulloh76/user-service/pkg/domain"
	"github.com/abdulloh76/user-service/pkg/handlers"
	"github.com/abdulloh76/user-service/pkg/handlers/userGrpc"
	"github.com/abdulloh76/user-service/pkg/store"
	"github.com/abdulloh76/user-service/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	configs := utils.LoadConfig("./", "dev", "env")

	dynamoDB := store.NewDynamoDBStore(context.TODO(), configs.DYNAMO_TABLE_NAME)
	domain := domain.NewUsersDomain(dynamoDB)
	grpcServer := handlers.NewGRPCServer(domain)

	s := grpc.NewServer()
	userGrpc.RegisterUserServer(s, grpcServer)

	l, err := net.Listen("tcp", ":"+configs.PORT)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
