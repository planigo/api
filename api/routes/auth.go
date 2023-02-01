package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
	"planigo/pkg/auth"
)

func AuthRoutes(app fiber.Router, handler *auth.Handler) {
	router := app.Group("/auth")

	router.Post("/login", handler.Login())
	router.Get("/me", requireLogin(handler.Session), handler.Me())
	router.Get("/logout", requireLogin(handler.Session), handler.Logout())
}

func requireLogin(r *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := r.Get(c)
		if err != nil {
			panic(err)
		}
		if sess.Get("uid") == nil {
			return c.SendStatus(http.StatusUnauthorized)
		}
		return c.Next()
	}
}
