package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"planigo/core/auth"
	"planigo/core/presenter"
	"planigo/internal/entities"
	"planigo/pkg/mail"
	"planigo/pkg/store"
	"planigo/utils"
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
			return presenter.Error(ctx, fiber.StatusBadRequest, errors.New(presenter.WrongCredential))
		}

		findedUser, err := r.UserStore.FindUserByEmail(user.Email)
		if err != nil {
			return presenter.Error(ctx, fiber.StatusBadRequest, errors.New(presenter.WrongCredential))
		}

		if isSamePassword := utils.CheckPasswordHash(user.Password, findedUser.Password); !isSamePassword {
			return presenter.Error(ctx, fiber.StatusBadRequest, errors.New(presenter.WrongCredential))
		}

		return presenter.Response(ctx, fiber.StatusOK, &fiber.Map{
			"access_token": auth.GenerateJWT(&auth.TokenPayload{Id: findedUser.Id, Role: findedUser.Role}),
		})
	}
}

func (r AuthHandler) Me() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Locals("userId").(string)
		userById, err := r.UserStore.FindUserById(id)
		if err != nil {
			return presenter.Error(ctx, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(ctx, fiber.StatusOK, userById)
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
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}
		return presenter.Response(c, fiber.StatusOK, "")
	}
}
