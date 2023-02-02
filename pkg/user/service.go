package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"planigo/config/mail"
	"planigo/config/store"
	"planigo/pkg/entities"
)

type Handler struct {
	*store.Store
	*mail.Mailer
	Session *session.Store
}

func NewHandler(store *store.Store, mailer *mail.Mailer, session *session.Store) *Handler {
	return &Handler{store, mailer, session}
}

func (r Handler) FindUsers() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := r.UserStore.FindUsers()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(http.StatusCreated).JSON(users)
	}
}

func (r Handler) RegisterUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userPayload := new(entities.User)

		if err := ctx.BodyParser(&userPayload); err != nil {
			log.Fatal(err)
		}

		password, err := HashPassword(userPayload.Password)
		if err != nil {
			log.Fatal(err)
		}

		userPayload.Password = password

		if _, err = r.CreateUser(*userPayload); err != nil {
			log.Fatal(err)
		}

		sendValidationEmail(r.Mailer, *userPayload)

		return ctx.SendStatus(http.StatusCreated)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func sendValidationEmail(mailer *mail.Mailer, user entities.User) {
	emailContent := mail.Content{
		To:      user.Email,
		Subject: "Bienvenue sur Planigo",
		Body:    "Veuillez valider votre compte sur ce lien pour pouvoir effectuer votre premier rendez-vous: https://planigo.fr/",
	}

	mailer.Send(emailContent)
}
