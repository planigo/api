package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func RequireRoles(roles []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("userRole")

		authorized := false
		for _, role := range roles {
			if userRole == role {
				authorized = true
				break
			}
		}

		if !authorized {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"statusCode": "403",
				"message":    "You are not authorized to access this resource",
			})
		}

		return c.Next()
	}
}
