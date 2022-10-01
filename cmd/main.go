package main

import (
	"log"
	"net"
	"os"

	"github.com/abdulloh76/user-service/handlers"
	"github.com/abdulloh76/user-service/handlers/userGrpc"
	"google.golang.org/grpc"
)

func main() {
	// DYNAMODB_PORT := os.Getenv("DYNAMODB_PORT")
	PORT := os.Getenv("PORT")
	// tableName := "users"

	// router := gin.Default()

	// dynamoDB := store.NewDynamoDBStore(context.TODO(), DYNAMODB_PORT, tableName)
	// domain := domain.NewUsersDomain(dynamoDB)
	// handler := handlers.NewGinAPIHandler(domain)

	// handlers.RegisterHandlers(router, handler)

	// router.Run(":" + PORT)

	s := grpc.NewServer()
	srv := &handlers.GRPCServer{}

	userGrpc.RegisterGetServer(s, srv)

	l, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
