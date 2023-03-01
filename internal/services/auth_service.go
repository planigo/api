package services

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"planigo/core/auth"
	"planigo/internal/entities"
	"planigo/pkg/mail"
	"planigo/pkg/store"
)

type AuthHandler struct {
	*store.Store
	*mail.Mailer
}

func (r AuthHandler) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		user := &entities.User{}
		if err := ctx.BodyParser(user); err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		if user.Email == "" || user.Password == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Wrong email or password!",
			})
		}

		findedUser, err := r.UserStore.FindUserByEmail(user.Email)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Wrong email or password!",
			})
		}

		if isSamePassword := CheckPasswordHash(user.Password, findedUser.Password); !isSamePassword {
			return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Wrong email or password!",
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"access_token": auth.GenerateJWT(&auth.TokenPayload{ID: findedUser.Id, Role: findedUser.Role}),
		})
	}
}

func (r AuthHandler) Me() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Locals("userId").(string)
		userById, err := r.UserStore.FindUserById(id)
		if err != nil {
			return err
		}

		return ctx.Status(http.StatusOK).JSON(userById)
	}
}

func (r AuthHandler) Logout() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	}
}

func (r AuthHandler) ValidateEmail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Params("token")
		if err := r.UserStore.ValidateUserEmail(token); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"statusCode": http.StatusInternalServerError,
				"message":    "Something went wrong.",
			})
		}
		return c.SendStatus(http.StatusOK)
	}
}
