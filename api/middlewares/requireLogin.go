package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
)

func IsLoggedIn(r *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := r.Get(c)
		if err != nil {
			panic(err)
		}
		if sess.Get(sess.ID()) == nil {
			return c.SendStatus(http.StatusUnauthorized)
		}

		c.Locals("userId", sess.Get(sess.ID()))
		c.Locals("userRole", sess.Get("role"))

		return c.Next()
	}
}
