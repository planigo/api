package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/api/middlewares"
	"planigo/pkg/user"
)

func UserRoutes(app fiber.Router, handler *user.Handler) {
	router := app.Group("/users")

	router.Get("/",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin"}),
		handler.FindUsers(),
	)

	router.Post("/", handler.RegisterUser())
}
