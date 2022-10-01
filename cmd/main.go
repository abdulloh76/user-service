package main

import (
	"context"
	"os"

	"github.com/abdulloh76/user-service/domain"
	"github.com/abdulloh76/user-service/handlers"
	"github.com/abdulloh76/user-service/store"

	"github.com/gin-gonic/gin"
)

func main() {
	DYNAMODB_PORT := os.Getenv("DYNAMODB_PORT")
	PORT := os.Getenv("PORT")
	tableName := "users"

	router := gin.Default()

	dynamoDB := store.NewDynamoDBStore(context.TODO(), DYNAMODB_PORT, tableName)
	domain := domain.NewUsersDomain(dynamoDB)
	handler := handlers.NewGinAPIHandler(domain)

	handlers.RegisterHandlers(router, handler)

	router.Run(":" + PORT)
}
