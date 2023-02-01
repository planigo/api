package user

import "github.com/gofiber/fiber/v2"

func returnUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    []any{},
	})
}
