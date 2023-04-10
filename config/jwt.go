package settings

import (
	"elprogramador.co/go/echo/structs"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
)

func JsonWebTokenConfig() echojwt.Config {
	// JWT configuration
	keyBase64 := os.Getenv("PUBLIC_KEY")
	keyData, _ := base64.StdEncoding.DecodeString(keyBase64)
	key, _ := jwt.ParseRSAPublicKeyFromPEM(keyData)
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(structs.JwtCustomClaims)
		},
		SigningKey:    key,
		SigningMethod: "RS256",
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/api/v1/chat" {
				return true
			}
			return false
		},
	}
}
