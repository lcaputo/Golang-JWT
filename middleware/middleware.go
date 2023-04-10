package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartMiddlewares(e *echo.Echo) {
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	//e.Use(database.GormMiddleware)
	e.Use(middleware.CORS())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	//}))
}
