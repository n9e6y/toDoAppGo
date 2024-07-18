package services

import (
	"todoapp/models"
	"todoapp/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetUserByEmail(email string) (*models.User, error) {
	return repositories.GetUserByEmail(email)
}
