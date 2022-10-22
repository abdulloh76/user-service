package main

import (
	"context"
	"os"

	"github.com/abdulloh76/user-service/pkg/domain"
	"github.com/abdulloh76/user-service/pkg/handlers"
	"github.com/abdulloh76/user-service/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	PORT := os.Getenv("PORT")
	tableName := "users"

	router := gin.Default()

	dynamoDB := store.NewDynamoDBStore(context.TODO(), tableName)
	domain := domain.NewUsersDomain(dynamoDB)
	handler := handlers.NewGinAPIHandler(domain)

	handlers.RegisterHandlers(router, handler)

	router.Run(":" + PORT)
}
