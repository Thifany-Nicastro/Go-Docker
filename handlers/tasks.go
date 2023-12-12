package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Thifany-Nicastro/Go-Docker/database"
	"github.com/Thifany-Nicastro/Go-Docker/models"
)

func Home(c *fiber.Ctx) error {
    return c.SendString("Hi")
}

func CreateTask(c *fiber.Ctx) error {
    task := new(models.Task)

    if err := c.BodyParser(task); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    database.DB.Db.Create(&task)

    return c.Status(200).JSON(task)
}

func ListTasks(c *fiber.Ctx) error {
    tasks := []models.Task{}
	
    database.DB.Db.Find(&tasks)

    return c.Status(200).JSON(tasks)
}