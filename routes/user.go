package routes

import (
	"elprogramador.co/go/echo/handlers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserRoutes(router *echo.Group) {
	router.GET("/restricted", func(c echo.Context) error {
		return c.String(http.StatusOK, "You are in the restricted area!")
	})
	router.GET("/GetUsers", handlers.GetUsers)
	router.GET("/GetUserById/:id", handlers.GetUserById)
}
