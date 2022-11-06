package handlers

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/abdulloh76/user-service/pkg/domain"
	"github.com/abdulloh76/user-service/pkg/utils"
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
	router.GET("/:id", handler.GetHandler)
	router.PUT("/:id/credentials", handler.UpdateUserCredentialsHandler)
	router.PUT("/:id/password", handler.UpdatePasswordHandler)
	router.PUT("/:id/credentials", handler.UpdateAddressHandler)
	router.DELETE("/:id", handler.DeleteHandler)
}

func (g *GinAPIHandler) GetHandler(context *gin.Context) {
	id := context.Param("id")

	user, err := g.users.GetUser(context, id)

	if errors.Is(err, utils.ErrUserNotFound) {
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

func (g *GinAPIHandler) UpdatePasswordHandler(context *gin.Context) {
	id := context.Param("id")

	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = g.users.UpdatePassword(context, id, body)
	if errors.Is(err, utils.ErrUserNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if errors.Is(err, utils.ErrJsonUnmarshal) {
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

	context.JSON(http.StatusOK, map[string]string{
		"message": "success", // todo double check
	})
}

func (g *GinAPIHandler) UpdateAddressHandler(context *gin.Context) {
	id := context.Param("id")

	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = g.users.UpdateAddress(context, id, body)
	if errors.Is(err, utils.ErrUserNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if errors.Is(err, utils.ErrJsonUnmarshal) {
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
	context.JSON(http.StatusOK, map[string]string{
		"message": "success", // todo double check
	})
}

func (g *GinAPIHandler) UpdateUserCredentialsHandler(context *gin.Context) {
	id := context.Param("id")

	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = g.users.UpdateUserCredentials(context, id, body)
	if errors.Is(err, utils.ErrUserNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if errors.Is(err, utils.ErrJsonUnmarshal) {
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

	context.JSON(http.StatusOK, map[string]string{
		"message": "success", // todo double check
	})
}

func (g *GinAPIHandler) DeleteHandler(context *gin.Context) {
	id := context.Param("id")

	err := g.users.DeleteUser(context, id)
	if errors.Is(err, utils.ErrUserNotFound) {
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
