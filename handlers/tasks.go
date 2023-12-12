package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Thifany-Nicastro/Go-Docker/models"
	"github.com/Thifany-Nicastro/Go-Docker/repositories"
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

    repositories.CreateTask(task)

    return c.Status(200).JSON(task)
}

func ListTasks(c *fiber.Ctx) error {
    tasks := repositories.FindAllTasks()

    return c.Status(200).JSON(tasks)
}