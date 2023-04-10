package handlers

import (
	"elprogramador.co/go/echo/api"
	"elprogramador.co/go/echo/database"
	"elprogramador.co/go/echo/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUsers(c echo.Context) error {
	// get all users
	var users []models.User
	//DB := database.Conn(c.Request().Context())
	database.DB.Find(&users)

	return api.SendResponse(c, http.StatusOK, "List of all users", users)
}

func GetUserById(c echo.Context) error {
	//type params struct {
	//	Id string `json:"id" validate:"required"`
	//}
	//var idUser = params{
	//	Id: c.Param("id"),
	//}
	//if err := c.Validate(idUser); err != nil {
	//	return err
	//}
	user := new(models.User)
	database.DB.First(&user, c.Param("id"))
	if user.ID == 0 {
		return api.SendResponse(c, http.StatusNotFound, "User not found", nil)
	}
	return api.SendResponse(c, http.StatusOK, "User", user)
}
