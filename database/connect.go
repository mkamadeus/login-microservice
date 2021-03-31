package database

import (
	"fmt"
	"login-microservice/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to database
func ConnectToDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=password dbname=logintest port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if(err != nil) {
		fmt.Println("Error when connecting to database")
	}

	DB.AutoMigrate(&model.User{})
	fmt.Println("Database migrated")
}