package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/teerapon19/todos/env"
	"github.com/teerapon19/todos/mongodb"
	"github.com/teerapon19/todos/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env.Load()

	app := fiber.New()
	app.Use(cors.New())

	mongodb.Connect()

	router.SetRouter(app)
	log.Fatal(app.Listen(":8080"))
}
