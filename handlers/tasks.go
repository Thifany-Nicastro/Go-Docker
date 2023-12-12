package handlers

import (
	"strconv"

	"github.com/Thifany-Nicastro/Go-Docker/models"
	"github.com/Thifany-Nicastro/Go-Docker/repositories"
	"github.com/gofiber/fiber/v2"
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

func DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Fail",
			"err":     err,
		})
	}

	repositories.DeleteTask(id)

	return c.Status(200).JSON("Task deleted")
}

func FindTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Fail",
			"err":     err,
		})
	}

	task := repositories.FindTaskById(id)

	return c.Status(200).JSON(task)
}
