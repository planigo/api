package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/internal/services"
)

func UserRoutes(app fiber.Router, handler *services.UserHandler) {
	router := app.Group("/users")

	router.Get("/", handler.FindUsers())

	router.Post("/", handler.RegisterUser())
}
