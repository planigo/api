package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/pkg/category"
)

func CategoryRoutes(app fiber.Router, handler *category.Handler) {
	categoryRoutes := app.Group("/categories")

	categoryRoutes.Get("/", handler.GetCategories())
}
