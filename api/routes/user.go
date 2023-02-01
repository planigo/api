package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/pkg/user"
)

func UserRoutes(app fiber.Router, handler *user.Handler) {
	router := app.Group("/users")

	router.Get("/", handler.FindUsers())
}
