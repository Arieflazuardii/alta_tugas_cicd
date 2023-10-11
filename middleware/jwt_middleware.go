package middleware

import (
	"praktikum/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}

func JwtMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey : []byte(constants.SECRET_KEY),
		SigningMethod : "HS256",
	})
}