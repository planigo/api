package auth

import (
	"fmt"
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

func (r Handler) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		userPayload := new(entities.User)
		if err := ctx.BodyParser(&userPayload); err != nil {
			log.Fatal(err)
		}

		user, err := r.UserStore.FindUserByEmail(userPayload.Email)
		if err != nil {
			log.Fatal(err)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPayload.Password)); err != nil {
			log.Fatal(err)
		}

		sess, err := r.Session.Get(ctx)
		if err != nil {
			panic(err)
		}

		sid := sess.ID()
		fmt.Println(sid)

		sess.Set("uid", sid)
		if err := sess.Save(); err != nil {
			panic(err)
		}

		return ctx.SendStatus(http.StatusOK)
	}
}

func (r Handler) Me() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sess, err := r.Session.Get(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println(sess)
		fmt.Println(sess.Get("uid"))

		return ctx.Status(http.StatusOK).JSON(fiber.Map{"isLoggedIn": true})
	}
}

func (r Handler) Logout() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sess, err := r.Session.Get(ctx)
		if err != nil {
			panic(err)
		}

		err = sess.Destroy()
		if err != nil {
			return err
		}

		return ctx.SendStatus(http.StatusOK)
	}
}
