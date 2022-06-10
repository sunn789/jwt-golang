package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sunn789/authInGo/user"
	"github.com/sunn789/authInGo/user/controllers"
)

func main() {
	e := echo.New()
	e.GET("/", CheckHealty)
	e.GET("/user/signin", controllers.SignInForm()).Name = "userSignInForm"
	e.POST("/user/signin", controllers.SignIn())

	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &user.Claims{},
		SigningKey:              []byte(user.GetJWTSecret()),
		TokenLookup:             "cookie:vin-cookie", // "<source>:<name>"
		ErrorHandlerWithContext: user.JWTErrorChecker,
	}))
	adminGroup.Use(user.TokenRefresherMiddleware)
	adminGroup.GET("", controllers.Admin())

	e.Logger.Fatal(e.Start(":3000"))
}
func CheckHealty(e echo.Context) error {
	return e.String(http.StatusOK, "Ok")
}
