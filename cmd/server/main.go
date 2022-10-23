package main

import (
	"context"

	"github.com/abdulloh76/user-service/pkg/domain"
	"github.com/abdulloh76/user-service/pkg/handlers"
	"github.com/abdulloh76/user-service/pkg/store"
	"github.com/abdulloh76/user-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	configs := utils.LoadConfig("./", "dev", "env")

	router := gin.Default()

	dynamoDB := store.NewDynamoDBStore(context.TODO(), configs.DYNAMO_TABLE_NAME)
	domain := domain.NewUsersDomain(dynamoDB)
	handler := handlers.NewGinAPIHandler(domain)

	handlers.RegisterHandlers(router, handler)

	router.Run(":" + configs.PORT)
}
