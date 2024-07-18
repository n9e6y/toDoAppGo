package db

import (
	"log"
	"todoapp/config"
	"todoapp/models"
)

func Migrate() {
	err := config.DB.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
