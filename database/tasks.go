package database

import (
	"log"

	"github.com/holanda-lucas/gotasks-back/models"
)

func RegisterTask(t *models.Task) uint{
	err := t.ValidateTask()

	if err != nil {
		log.Fatal("Cannot register task: " + err.Error())
	}

	DB.Create(&t)

	return t.ID
}

func GetTask(id uint) models.Task {
	var task models.Task

	DB.First(&task, id)
	return task
}

func EditTask(t *models.Task) {
	err := t.ValidateTask()

	if err != nil || GetTask(t.ID).Title == "" {
		log.Fatal("Cannot edit task: " + err.Error())
	}

	DB.Table("tasks").UpdateColumns(t)
}

func DeleteTask(id uint) {
	DB.Unscoped().Delete(&models.Task{}, id)
}

func GetTasksFromUser(user_id uint) []models.Task {
	var tasks []models.Task

	DB.Where(&models.Task{User_id: user_id}).Find(&tasks)
	return tasks
}