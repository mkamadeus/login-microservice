package main

import (
	"login-microservice/database"
	"login-microservice/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectToDatabase()
	router.SetupRoutes(app)

	app.Listen(":3000")
}