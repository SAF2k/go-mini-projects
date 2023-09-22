package database

import (
	"auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/auth"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
