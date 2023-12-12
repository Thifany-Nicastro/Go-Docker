package main

import (
	"github.com/Thifany-Nicastro/Go-Docker/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListTasks)

	app.Post("/tasks", handlers.CreateTask)

	app.Delete("/tasks/:id", handlers.DeleteTask)
}
