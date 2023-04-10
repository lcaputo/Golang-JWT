package main

import (
	"context"
	"elprogramador.co/go/echo/database"
	"elprogramador.co/go/echo/helpers"
	"elprogramador.co/go/echo/middleware"
	"elprogramador.co/go/echo/models"
	"elprogramador.co/go/echo/routes"
	"elprogramador.co/go/echo/structs"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"elprogramador.co/go/echo/api"
	config "elprogramador.co/go/echo/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return api.SendResponse(c, http.StatusOK, "Hello, World!", nil)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		var errors []structs.ApiRequestError
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			var el structs.ApiRequestError
			el.Field = helpers.PrintJSONTag(i, validationError.Field())
			el.Error = validationError.Tag()
			el.Message = validationError.Value()
			errors = append(errors, el)
		}
		response := structs.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid form",
			Data:    errors,
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
		//return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	if os.Getenv("GO_ENV") != "production" {
		_ = godotenv.Load(".env")
	}

	// Database
	database.DBConnection()
	if err := database.DB.AutoMigrate(models.User{}); err != nil {
		return
	}
	if err := database.DB.AutoMigrate(models.Chat{}); err != nil {
		return
	}

	//if err := security.LoadRSAKeys(); err != nil {
	//	log.Println("Err RSA Kes:", err)
	//	os.Exit(0)
	//}
	//log.Println("RSA KEYS OK!")
	//if err := database.PrepareConnection(); err != nil {
	//	log.Println("Err database:", err)
	//	os.Exit(0)
	//}
	//log.Println("Database connected")

	e := echo.New()
	e.HideBanner = true

	middleware.StartMiddlewares(e)

	routes.StartRoutes(e)

	e.Validator = &CustomValidator{validator: validator.New()}

	go func() {
		// Wait for an interrupt signal
		for sig := range c {
			// Try to stop the Echo server
			if err := e.Server.Shutdown(context.Background()); err != nil {
				// If an error occurs while stopping the server, print an error message and exit the program
				log.Println("error trying to stop echo service:", err)
				return
			}
			// If the server stops successfully, print a log message and exit the program
			log.Println("stopping echo service:", sig)
			return
		}
	}()

	err := e.Start(config.FindEnvOrDefault("HOST", "localhost:1323"))
	if err != nil {
		log.Println("Error staring echo service:", err)
		os.Exit(0)
	}

}
