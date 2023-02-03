package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/handlers"
)

func UserRoutes(app fiber.Router, handler *handlers.UserHandler) {
	router := app.Group("/users")

	router.Get("/", handler.FindUsers())

	router.Post("/", handler.RegisterUser())
}
