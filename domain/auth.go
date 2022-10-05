package domain

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"

	"github.com/abdulloh76/user-service/types"
	"github.com/abdulloh76/user-service/utils"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "auth_salt"
	signingKey = "signing_key"
	tokenTTL   = 48 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type AuthDomain struct {
	store types.AuthStore
}

func NewAuthDomain(d types.AuthStore) *AuthDomain {
	return &AuthDomain{
		store: d,
	}
}

func (a *AuthDomain) CreateUser(ctx context.Context, body []byte) (string, error) {
	userCredentials := types.UserCredentials{}
	if err := json.Unmarshal(body, &userCredentials); err != nil {
		return "", utils.ErrJsonUnmarshal
	}

	userCredentials.Password = generatePasswordHash(userCredentials.Password)
	return a.store.CreateUser(ctx, userCredentials)
}

func (a *AuthDomain) GenerateToken(ctx context.Context, body []byte) (string, error) {
	userCredentials := types.UserCredentials{}
	if err := json.Unmarshal(body, &userCredentials); err != nil {
		return "", utils.ErrJsonUnmarshal
	}

	user, err := a.store.GetUser(ctx, userCredentials.Email, generatePasswordHash(userCredentials.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (a *AuthDomain) ParseToken(body []byte) (string, error) {
	type tokenReqBody struct {
		token string
	}
	reqBody := tokenReqBody{}
	if err := json.Unmarshal(body, &reqBody); err != nil {
		return "", utils.ErrJsonUnmarshal
	}

	token, err := jwt.ParseWithClaims(reqBody.token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.ErrInvalidJWTMethod
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", utils.ErrInvalidTokenClaims
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
