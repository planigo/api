package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/core/middlewares"
	"planigo/internal/services"
)

func AuthRoutes(app fiber.Router, handler *services.AuthHandler) {
	router := app.Group("/auth")

	router.Post("/login", handler.Login())

	router.Get(
		"/me",
		middlewares.IsLoggedIn,
		handler.Me(),
	)

	router.Get(
		"/logout",
		middlewares.IsLoggedIn,
		handler.Logout(),
	)

	router.Get(
		"/validate/:token",
		handler.ValidateEmail(),
	)
}
