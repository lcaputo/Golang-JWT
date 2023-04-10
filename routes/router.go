package routes

import (
	config "elprogramador.co/go/echo/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func StartRoutes(e *echo.Echo) {
	AuthRoutes(e)
	EventRoutes(e)
	apiRoute := e.Group("/api/v1")
	apiRoute.Use(echojwt.WithConfig(config.JsonWebTokenConfig()))
	UserRoutes(apiRoute)
	ChatRoutes(apiRoute)
}
