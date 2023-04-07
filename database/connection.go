package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DATABASE_URL = "postgres://admin:password@localhost:5432/postgres?sslmode=disable"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})
	if error != nil {
		log.Fatal("failed to connect database")
	} else {
		log.Println("database connected")
	}
}
