package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teerapon19/todos/model"
	"github.com/teerapon19/todos/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ShowTasks(c *fiber.Ctx) error {

	tasks, err := mongodb.TaskGetAll()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "tasks": tasks})
}

func ShowSingleTask(c *fiber.Ctx) error {
	id := c.Params("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if id == "" || err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "ID is empty!"})
	}

	task, err := mongodb.TaskGetSingle(oid)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "task": task})
}

func MarkTaskAsAccomplished(c *fiber.Ctx) error {
	id := c.Params("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if id == "" || err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "ID is empty!"})
	}

	_, err = mongodb.TaskMarkAction(oid, true)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Wow! It look like someone very happy. 😋"})
}

func UnmarkTaskAsAccomplished(c *fiber.Ctx) error {
	id := c.Params("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if id == "" || err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "ID is empty!"})
	}

	_, err = mongodb.TaskMarkAction(oid, false)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Well?. 🤔"})
}

func CreateTask(c *fiber.Ctx) error {
	var task model.Task

	if err := c.BodyParser(&task); err != nil {
		return err
	}

	_, err := mongodb.TaskInsertNew(task)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Create success! Happy with your task. 😉"})
}

func UpdateEditTask(c *fiber.Ctx) error {
	var task model.Task

	if err := c.BodyParser(&task); err != nil {
		return err
	}

	_, err := mongodb.TaskUpdateEdited(task)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Update success! Happy with your task. 😁"})
}

func RemoveTask(c *fiber.Ctx) error {

	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if id == "" || err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "ID is empty!"})
	}

	_, err = mongodb.TaskDelete(oid)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Remove success! Play hard Work harder. 🤣"})
}
