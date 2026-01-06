package helper

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)


func GetAuthId(c echo.Context) int {
	user := c.Get("user")
	if user == nil {
		return 0
	}

	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	userIdStr, ok := claims["user_id"].(string)
	if !ok {
		return 0
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0
	}

	return userId
}



