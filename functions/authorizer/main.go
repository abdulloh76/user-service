package main

import (
	"context"
	"os"

	"github.com/abdulloh76/user-service/pkg/domain"
	"github.com/abdulloh76/user-service/pkg/handlers"
	"github.com/abdulloh76/user-service/pkg/store"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	DYNAMO_TABLE_NAME := os.Getenv("DYNAMO_TABLE_NAME")

	dynamoDB := store.NewAuthDynamoStore(context.TODO(), DYNAMO_TABLE_NAME)
	domain := domain.NewAuthDomain(dynamoDB)
	handler := handlers.NewAuthApiHandler(domain)

	lambda.Start(handler.AuthMiddleware)
}
