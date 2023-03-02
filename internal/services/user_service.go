package services

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"planigo/core/auth"
	"planigo/core/enums"
	"planigo/core/presenter"
	"planigo/internal/entities"
	"planigo/pkg/mail"
	"planigo/pkg/store"
	"planigo/utils"
)

type UserHandler struct {
	*store.Store
	*mail.Mailer
}

func (r UserHandler) FindUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := r.UserStore.FindUsers()
		if err != nil {
			log.Fatal(err)
		}

		return presenter.Response(c, fiber.StatusOK, users)
	}
}

func (r UserHandler) RegisterUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		userPayload := ParseUserBody(c)

		password := utils.HashPassword(userPayload.Password)
		userPayload.Password = password

		uuid, err := r.CreateUser(*userPayload)
		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}

		userPayload.Id = uuid

		if err := sendValidationEmail(r.Mailer, userPayload); err != nil {
			return err
		}

		return presenter.Response(c, fiber.StatusCreated, userPayload)
	}
}

func (r UserHandler) RegisterCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		registerCustomerBody := struct {
			LastName        string `json:"lastname"`
			FirstName       string `json:"firstname"`
			Email           string `json:"email"`
			Password        string `json:"password"`
			ConfirmPassword string `json:"confirmPassword"`
		}{}

		if err := c.BodyParser(&registerCustomerBody); err != nil {
			return presenter.Error(c, fiber.StatusBadRequest, err)
		}
		if registerCustomerBody.Password != registerCustomerBody.ConfirmPassword {
			return presenter.Error(c, fiber.StatusBadRequest, errors.New(presenter.PasswordsNotMatch))
		}

		userPayload := &entities.User{
			Lastname:  registerCustomerBody.LastName,
			Firstname: registerCustomerBody.FirstName,
			Email:     registerCustomerBody.Email,
			Role:      enums.Customer,
		}

		password := utils.HashPassword(userPayload.Password)
		userPayload.Password = password

		uuid, err := r.CreateUser(*userPayload)
		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}

		userPayload.Id = uuid

		if err := sendValidationEmail(r.Mailer, userPayload); err != nil {
			fmt.Println("Error while sending email: ", err.Error())
		}

		return presenter.Response(c, fiber.StatusCreated, userPayload)
	}
}

func ParseUserBody(c *fiber.Ctx) *entities.User {
	userPayload := new(entities.User)

	if err := c.BodyParser(&userPayload); err != nil {
		log.Fatal(err)
	}

	return userPayload
}

func sendValidationEmail(mailer *mail.Mailer, user *entities.User) error {
	validationToken := auth.GenerateJWT(&auth.TokenPayload{Id: user.Id, Role: user.Role})
	validateUrl := os.Getenv("FRONTEND_URL") + "/validate/" + validationToken
	emailContent := mail.Content{
		To:      user.Email,
		Subject: "Bienvenue sur Planigo",
		Body:    "Bienvenue sur Planigo, votre application de prise de reservation en ligne. Merci de cliquer sur le lien suivant pour valider votre compte : " + validateUrl,
	}

	if err := mailer.Send(emailContent); err != nil {
		return err
	}
	return nil
}
