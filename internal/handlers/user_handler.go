package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/core/middlewares"
	"planigo/internal/services"
)

func UserRoutes(app fiber.Router, handler *services.UserHandler) {
	router := app.Group("/users")

	router.Get("/",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{"admin"}),
		handler.FindUsers(),
	)

	router.Post("/", handler.RegisterUser())
	router.Post("/customer", handler.RegisterCustomer())
}
