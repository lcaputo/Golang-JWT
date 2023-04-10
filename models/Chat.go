package models

import "gorm.io/gorm"

type Chat struct {
	*gorm.Model

	// user id foreign key to user table
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	User     User   `json:"user" gorm:"foreignkey:ID"`
	Message  string `json:"message"`
}
