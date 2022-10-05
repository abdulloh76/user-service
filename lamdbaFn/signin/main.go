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
	DYNAMODB_PORT := os.Getenv("DYNAMODB_PORT")
	tableName := "users"

	dynamoDB := store.NewAuthDynamoStore(context.TODO(), DYNAMODB_PORT, tableName)
	domain := domain.NewAuthDomain(dynamoDB)
	handler := handlers.NewAuthApiHandler(domain)

	lambda.Start(handler.SignIn)
}
