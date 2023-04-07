package models

import (
	"gorm.io/gorm"
)

type UserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model

	Name     string `json:"name"`
	Email    string `json:"email" gorm:"primaryKey" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
