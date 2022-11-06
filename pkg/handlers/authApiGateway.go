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
		return utils.ErrResponse(http.StatusBadRequest, utils.ErrEmptyBody.Error()), utils.ErrEmptyBody
	}

	id, err := g.auth.CreateUser(ctx, []byte(event.Body))
	if err != nil {
		if errors.Is(err, utils.ErrJsonUnmarshal) {
			return utils.ErrResponse(http.StatusBadRequest, err.Error()), err
		}
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), err
	}

	return utils.Response(http.StatusOK, map[string]string{"id": id}), nil
}

func (g *AuthApiHandler) SignIn(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if strings.TrimSpace(event.Body) == "" {
		return utils.ErrResponse(http.StatusBadRequest, utils.ErrEmptyBody.Error()), utils.ErrEmptyBody
	}

	token, err := g.auth.GenerateToken(ctx, []byte(event.Body))
	if err != nil {
		if errors.Is(err, utils.ErrWithDB) {
			return utils.ErrResponse(http.StatusInternalServerError, err.Error()), err
		}
		return utils.ErrResponse(http.StatusBadRequest, err.Error()), err
	}

	return utils.Response(http.StatusOK, map[string]string{"token": token}), nil
}

func (g *AuthApiHandler) AuthMiddleware(ctx context.Context, event events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	if strings.TrimSpace(event.Headers["authorization"]) == "" {
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: false}, errors.New("token not provided")
	}

	userId, err := g.auth.ParseToken(event.Headers["authorization"])
	if err != nil {
		if errors.Is(err, utils.ErrInvalidJWTMethod) || errors.Is(err, utils.ErrInvalidTokenClaims) {
			return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: false}, err
		}
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: false}, errors.New("token not valid")
	}

	return events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: true,
		Context: map[string]interface{}{
			"userId": userId,
		}}, nil
}
