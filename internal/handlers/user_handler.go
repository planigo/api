package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/core/middlewares"
	"planigo/internal/services"
)

func UserRoutes(app fiber.Router, handler *services.UserHandler) {
	router := app.Group("/users")

	router.Get("/",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin"}),
		handler.FindUsers(),
	)

	router.Post("/", handler.RegisterUser())
}
