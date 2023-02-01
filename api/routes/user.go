package routes

import (
	"github.com/gofiber/fiber/v2"
	service "planigo/pkg/user"
)

func UserRoutes(app fiber.Router) {
	router := app.Group("/users")

	router.Get("/", service.GetUsers)
}
