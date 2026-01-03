package helper

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginTokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

func Login(userId, role int, name, email string, ) (*LoginTokenResponse, error) {
	expiredTime := time.Now().Add(24 * time.Hour)

	claims := JwtCustomClaims{
		ID:    strconv.Itoa(userId),
		Name:  name,
		Email: email,
		Role:  strconv.Itoa(role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := NewTokenUseCase().GenerateAccessToken(claims)
	if err != nil {
		return nil, err
	}

	return &LoginTokenResponse{
		Token:     token,
		ExpiresAt: expiredTime.Format(time.RFC3339),
	}, nil
}
