package structs

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
