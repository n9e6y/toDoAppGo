package repositories

import (
	"todoapp/config"
	"todoapp/models"
)

func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Find(&tasks).Error
	return tasks, err
}

func GetTask(id uint) (*models.Task, error) {
	var task models.Task
	err := config.DB.First(&task, id).Error
	return &task, err
}

func UpdateTask(task *models.Task) error {
	return config.DB.Save(task).Error
}

func DeleteTask(id uint) error {
	return config.DB.Delete(&models.Task{}, id).Error
}
