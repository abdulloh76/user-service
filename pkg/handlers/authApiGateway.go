package handlers

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/abdulloh76/user-service/pkg/domain"
	"github.com/abdulloh76/user-service/pkg/utils"
	"github.com/aws/aws-lambda-go/events"
)

type AuthApiHandler struct {
	auth *domain.AuthDomain
}

func NewAuthApiHandler(d *domain.AuthDomain) *AuthApiHandler {
	return &AuthApiHandler{auth: d}
}

func (g *AuthApiHandler) SignUp(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if strings.TrimSpace(event.Body) == "" {
		return utils.ErrResponse(http.StatusBadRequest, "empty request body"), nil
	}

	id, err := g.auth.CreateUser(ctx, []byte(event.Body))
	if err != nil {
		if errors.Is(err, utils.ErrJsonUnmarshal) {
			return utils.ErrResponse(http.StatusBadRequest, err.Error()), nil
		}
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return utils.Response(http.StatusOK, map[string]string{"id": id}), nil
}

func (g *AuthApiHandler) SignIn(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if strings.TrimSpace(event.Body) == "" {
		return utils.ErrResponse(http.StatusBadRequest, "empty request body"), nil
	}

	token, err := g.auth.GenerateToken(ctx, []byte(event.Body))
	if err != nil {
		if errors.Is(err, utils.ErrWithDB) {
			return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
		}
		return utils.ErrResponse(http.StatusBadRequest, err.Error()), nil
	}

	return utils.Response(http.StatusOK, map[string]string{"token": token}), nil
}

func (g *AuthApiHandler) AuthMiddleware(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if strings.TrimSpace(event.Headers["token"]) == "" {
		return utils.ErrResponse(http.StatusBadRequest, "there is no token in headers"), nil
	}

	userId, err := g.auth.ParseToken(event.Headers["token"])
	if err != nil {
		if errors.Is(err, utils.ErrInvalidJWTMethod) || errors.Is(err, utils.ErrInvalidTokenClaims) {
			return utils.ErrResponse(http.StatusBadRequest, err.Error()), nil
		}
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return utils.Response(http.StatusOK, map[string]string{"userId": userId}), nil
}
