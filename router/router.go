package router

import (
	"login-microservice/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Setup routers for API
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/user/:id", services.GetUserByID)
	api.Post("/user/", services.CreateUser)
}