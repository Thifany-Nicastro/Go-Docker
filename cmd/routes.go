package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/Thifany-Nicastro/Go-Docker/handlers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/", handlers.ListTasks)

	app.Post("/tasks", handlers.CreateTask)
}