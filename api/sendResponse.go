package api

import (
	"elprogramador.co/go/echo/structs"
	"github.com/labstack/echo/v4"
)

// func SendResponse(c echo.Context, status int, message string, data ...interface{}) error {
// 	return c.JSON(status, Response{
// 		Status:  status,
// 		Message: message,
// 		Data:    data,
// 	})
// }

func SendResponse(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, structs.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
