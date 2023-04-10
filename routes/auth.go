package routes

import (
	"elprogramador.co/go/echo/api"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	e.POST("/login", api.Login)
	e.POST("/register", api.Register)
	e.GET("/tokenInfo", api.TokenInfo)
	//e.POST("/logout", api.Logout)
}
