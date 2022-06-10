package user

import (
	//"time"

	"github.com/golang-jwt/jwt"
	//"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string
	Admin string
	jwt.StandardClaims
}

// func login(e echo.Context) error {
// 	username := e.FormValue("username")
// 	paswword := e.FormValue("password")

// 	if username != "sunn789" || paswword != "123" {
// 		return echo.ErrUnauthorized
// 	}
// 	claims:= &jwtCustomClaims{
// 		"Ariel Sunny",
// 	"true",
// 	jwt.StandardClaims{
// 		ExpiresAt: time.Now().Add(time.Hour*72).Unix(),
// 	},
// 	}
// 	claims = nil
// 	return nil
// }
