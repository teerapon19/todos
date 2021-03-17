package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/teerapon19/todos/handler"
)

func SetRouter(app *fiber.App) {

	api := app.Group("/", logger.New())
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello! This is todos RESTful API. 👏"})
	})

	task := api.Group("/task")
	task.Get("/all", handler.ShowTasks)
	task.Get("/:id", handler.ShowSingleTask)
	task.Put("/mark/:id", handler.MarkTaskAsAccomplished)
	task.Put("/unmark/:id", handler.UnmarkTaskAsAccomplished)
	task.Post("/", handler.CreateTask)
	task.Put("/", handler.UpdateEditTask)
	task.Delete("/:id", handler.RemoveTask)
}
