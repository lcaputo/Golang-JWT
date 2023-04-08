package structs

import "github.com/golang-jwt/jwt/v4"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ApiRequestError struct {
	Field   string `json:"field"`
	Error   string `json:"error"`
	Message any    `json:"message"`
}

type AuthRouterDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
