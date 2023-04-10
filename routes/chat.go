package routes

import (
	"elprogramador.co/go/echo/api"
	"elprogramador.co/go/echo/database"
	"elprogramador.co/go/echo/events"
	"elprogramador.co/go/echo/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

func ChatRoutes(router *echo.Group) {
	handlerEvents := events.NewHandlerEvent()

	router.GET("/chat", handlerEvents.Handler)

	router.GET("/GetChat", func(c echo.Context) error {
		var messages []models.Chat
		database.DB.Find(&messages)

		return api.SendResponse(c, http.StatusOK, "Token info", messages)
	})

	router.POST("/SendMessage", func(c echo.Context) error {
		// obtener el token del request

		tokenString := c.Request().Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Get the key from the public key file
			keyBase64 := os.Getenv("PUBLIC_KEY")
			keyData, _ := base64.StdEncoding.DecodeString(keyBase64)
			publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(keyData)

			return publicKey, nil
		})

		// Check if there was an error parsing the token
		if err != nil {
			fmt.Printf("error parsing token: %v\n", err)
			return err
		}

		// Check if the token is valid
		if !token.Valid {
			fmt.Printf("token is not valid\n")
			return err
		}

		// Get the claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Printf("error getting claims\n")
			return err
		}

		user := models.User{}
		database.DB.Where("id = ?", claims["id"]).First(&user)

		message := models.Chat{
			Message:  c.FormValue("message"),
			UserID:   int64(user.ID),
			UserName: user.Name,
			User:     user,
		}

		err = c.Bind(&message)
		if err != nil {
			return err
		}
		database.DB.Create(&message)

		data, _ := json.Marshal(message)
		handlerEvents.Broadcast(events.EventMessage{
			EventName: "NewMessage",
			Data:      string(data),
		})

		return api.SendResponse(c, http.StatusOK, "Token info", message)
	})
}
