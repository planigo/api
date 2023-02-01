package routes

import "github.com/gofiber/fiber/v2"

func GetUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello, World2!",
		"status":  "success",
	})
}
