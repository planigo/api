package services

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"planigo/core/auth"
	"planigo/core/presenter"
	"planigo/internal/entities"
	"planigo/pkg/mail"
	"planigo/pkg/store"
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

		password := HashPassword(userPayload.Password)
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

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func ParseUserBody(c *fiber.Ctx) *entities.User {
	userPayload := new(entities.User)

	if err := c.BodyParser(&userPayload); err != nil {
		log.Fatal(err)
	}

	return userPayload
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func sendValidationEmail(mailer *mail.Mailer, user *entities.User) error {
	validationToken := auth.GenerateJWT(&auth.TokenPayload{ID: user.Id, Role: user.Role})
	emailContent := mail.Content{
		To:      user.Email,
		Subject: "Bienvenue sur Planigo",
		Body:    "Bienvenue sur Planigo, votre application de prise de reservation en ligne. Merci de cliquer sur le lien suivant pour valider votre compte: http://localhost:3000/validate?token=" + validationToken,
	}

	if err := mailer.Send(emailContent); err != nil {
		return err
	}
	return nil
}
