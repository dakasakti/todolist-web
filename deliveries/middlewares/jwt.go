package middlewares

import (
	"time"

	"github.com/dakasakti/todolist-web/constants"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTSign() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(constants.SECRET_JWT),
		SigningMethod: jwt.SigningMethodHS256.Name,
	})
}

func CreateToken(id uint, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["email"] = email
	claims["expired"] = time.Now().Add(time.Hour * 48).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(float64)
		return userId
	}

	return 0
}

func ExtractToken(value string) float64 {
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SECRET_JWT), nil
	})

	if token.Valid {
		userId := claims["id"].(float64)
		return userId
	}

	return 0
}
