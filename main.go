package main

import (
	"elprogramador.co/go/echo/helpers"
	"elprogramador.co/go/echo/structs"
	"encoding/base64"
	echojwt "github.com/labstack/echo-jwt/v4"
	"net/http"
	"os"

	"elprogramador.co/go/echo/api"
	"elprogramador.co/go/echo/database"
	"elprogramador.co/go/echo/models"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
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
	e := echo.New()

	// Database
	database.DBConnection()
	err := database.DB.AutoMigrate(models.User{})
	if err != nil {
		return
	}

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// Configurar el middleware JWT

	e.Validator = &CustomValidator{validator: validator.New()}

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/register", api.Register)

	type UserDTO struct {
		Email string `json:"email"`
	}
	// e.Validator = &api.CustomValidator{validator: validator.New()}

	// TODO: Activate login and add !Logout
	e.POST("/login", api.Login)

	// JWT configuration
	keyBase64 := os.Getenv("PUBLIC_KEY")
	keyData, _ := base64.StdEncoding.DecodeString(keyBase64)
	key, _ := jwt.ParseRSAPublicKeyFromPEM(keyData)
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(structs.JwtCustomClaims)
		},
		SigningKey:    key,
		SigningMethod: "RS256",
	}

	// Restricted group with JWT middleware
	apiRoute := e.Group("/api/v1")
	apiRoute.Use(echojwt.WithConfig(config))
	apiRoute.GET("/restricted", func(c echo.Context) error {
		return c.String(http.StatusOK, "You are in the restricted area!")
	})

	// Configure middleware with the custom claims type

	//r.Use(echojwt.WithConfig(config))
	//r.GET("", handlers.Restricted)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
