package services

import (
	"todoapp/models"
	"todoapp/repositories"
)

func CreateTask(task *models.Task) error {
	return repositories.CreateTask(task)
}

func GetTasks() ([]models.Task, error) {
	return repositories.GetTasks()
}

func GetTask(id uint) (*models.Task, error) {
	return repositories.GetTask(id)
}

func UpdateTask(task *models.Task) error {
	return repositories.UpdateTask(task)
}

func DeleteTask(id uint) error {
	return repositories.DeleteTask(id)
}
