package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"planigo/core/auth"
	"planigo/internal/entities"
	"planigo/pkg/mail"
	"planigo/pkg/store"
)

type UserHandler struct {
	*store.Store
	*mail.Mailer
	Session *session.Store
}

func NewUserHandler(store *store.Store, mailer *mail.Mailer, session *session.Store) *UserHandler {
	return &UserHandler{store, mailer, session}
}

func (r UserHandler) FindUsers() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := r.UserStore.FindUsers()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(http.StatusCreated).JSON(users)
	}
}

func (r UserHandler) RegisterUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		userPayload := ParseUserBody(ctx)

		password := HashPassword(userPayload.Password)
		userPayload.Password = password

		uuid, err := r.CreateUser(*userPayload)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		userPayload.Id = uuid

		if err := sendValidationEmail(r.Mailer, userPayload); err != nil {
			return err
		}

		return ctx.SendStatus(http.StatusCreated)
	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func ParseUserBody(ctx *fiber.Ctx) *entities.User {
	userPayload := new(entities.User)

	if err := ctx.BodyParser(&userPayload); err != nil {
		log.Fatal(err)
	}

	return userPayload
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func sendValidationEmail(mailer *mail.Mailer, user *entities.User) error {
	validationToken := auth.GenerateJWT(&auth.TokenPayload{ID: user.Id})
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
