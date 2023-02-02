package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/api/middlewares"

	"planigo/pkg/auth"
)

func AuthRoutes(app fiber.Router, handler *auth.Handler) {
	router := app.Group("/auth")

	router.Post("/login", handler.Login())

	router.Get(
		"/me",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin"}),
		handler.Me(),
	)

	router.Get(
		"/logout",
		middlewares.IsLoggedIn(handler.Session),
		handler.Logout(),
	)

	router.Get(
		"/validate/:token",
		handler.ValidateEmail(),
	)
}
