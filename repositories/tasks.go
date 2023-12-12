package repositories

import (
	"github.com/Thifany-Nicastro/Go-Docker/database"
	"github.com/Thifany-Nicastro/Go-Docker/models"
)

func FindAllTasks() []models.Task {
	tasks := []models.Task{}

	database.DB.Db.Find(&tasks)

	return tasks
}

func CreateTask(task *models.Task) {
	database.DB.Db.Create(task)
}

func DeleteTask(id int) {
	database.DB.Db.Delete(id)
}
