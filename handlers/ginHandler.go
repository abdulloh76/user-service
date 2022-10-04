package handlers

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/abdulloh76/user-service/domain"
	"github.com/gin-gonic/gin"
)

type GinAPIHandler struct {
	users *domain.Users
}

func NewGinAPIHandler(d *domain.Users) *GinAPIHandler {
	return &GinAPIHandler{
		users: d,
	}
}

func RegisterHandlers(router *gin.Engine, handler *GinAPIHandler) {
	router.POST("/user", handler.CreateHandler)
	router.GET("/user", handler.AllHandler)
	router.GET("/user/:id", handler.GetHandler)
	router.PUT("/user/:id", handler.PutHandler)
	router.DELETE("/user/:id", handler.DeleteHandler)
}

func (g *GinAPIHandler) AllHandler(context *gin.Context) {
	allUsers, err := g.users.AllUsers(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, allUsers)
}

func (g *GinAPIHandler) GetHandler(context *gin.Context) {
	id := context.Param("id")

	user, err := g.users.GetUser(context, id)

	if errors.Is(err, domain.ErrUserNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (g *GinAPIHandler) CreateHandler(context *gin.Context) {
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	newUser, err := g.users.CreateUser(context, body)
	if errors.Is(err, domain.ErrJsonUnmarshal) {
		context.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, newUser)
}

func (g *GinAPIHandler) PutHandler(context *gin.Context) {
	id := context.Param("id")

	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	updatedUser, err := g.users.ModifyUser(context, id, body)
	if errors.Is(err, domain.ErrUserNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if errors.Is(err, domain.ErrJsonUnmarshal) {
		context.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, updatedUser)
}

func (g *GinAPIHandler) DeleteHandler(context *gin.Context) {
	id := context.Param("id")

	err := g.users.DeleteUser(context, id)
	if errors.Is(err, domain.ErrUserNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted",
	})
}
