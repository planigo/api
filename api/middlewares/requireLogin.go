package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
)

func RequireLogin(r *session.Store) func(c *fiber.Ctx) error {
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
