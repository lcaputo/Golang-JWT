package routes

import (
	"elprogramador.co/go/echo/events"
	"fmt"
	"github.com/labstack/echo/v4"
)

func EventRoutes(router *echo.Echo) {
	handlerEvents := events.NewHandlerEvent()

	router.GET("/notify", handlerEvents.Handler)

	router.GET("/clients", func(c echo.Context) error {
		fmt.Println(handlerEvents.ListClients())
		return c.JSON(200, handlerEvents.ListClients())
	})

	router.GET("/send", func(c echo.Context) error {
		handlerEvents.Broadcast(events.EventMessage{
			EventName: "saludar",
			Data:      "Hello World",
		})
		return nil
	})
}
