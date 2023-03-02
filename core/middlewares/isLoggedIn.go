package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"planigo/core/auth"
	"strings"
)

func IsLoggedIn(c *fiber.Ctx) error {
	authHeader := c.GetReqHeaders()["Authorization"]
	if authHeader == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	jwt := authToken[1]
	payload, err := auth.VerifyJWT(jwt)
	if err != nil {
		fmt.Println(err.Error())
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Locals("userId", payload.Id)
	c.Locals("userRole", payload.Role)

	return c.Next()
}
