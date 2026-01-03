package helper

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TokenUseCase interface {
	GenerateAccessToken(claims JwtCustomClaims) (string, error)
}

type TokenUseCaseImpl struct{}

func NewTokenUseCase() TokenUseCase {
	return &TokenUseCaseImpl{}
}

func (t *TokenUseCaseImpl) GenerateAccessToken(claims JwtCustomClaims) (string, error) {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return "", errors.New("SECRET_KEY is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}


func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(401, "Missing token")
			}

			// ambil token setelah "Bearer "
			if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
				return echo.NewHTTPError(401, "Invalid token format")
			}
			tokenString := authHeader[7:]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Unexpected signing method")
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(401, "Invalid token")
			}

			c.Set("user", token)
			return next(c)
		}
	}
}
