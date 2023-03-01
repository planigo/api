package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/internal/services"
)

func CategoryRoutes(app fiber.Router, handler *services.CategoryHandler) {
	categoryRoutes := app.Group("/categories")

	categoryRoutes.Get("/", handler.GetCategories())
}
