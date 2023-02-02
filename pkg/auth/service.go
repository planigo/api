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

		user := &entities.User{}
		if err := ctx.BodyParser(user); err != nil {
			return err
		}

		if user.Email == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Email is required.")
		}

		if user.Password == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Password is required.")
		}

		findedUser, err := r.UserStore.FindUserByEmail(user.Email)
		if err != nil {
			log.Fatal(err)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(findedUser.Password), []byte(user.Password)); err != nil {
			log.Fatal(err)
		}

		sess, err := r.Session.Get(ctx)
		if err != nil {
			panic(err)
		}

		sid := sess.ID()
		sess.Set(sid, findedUser.Id)
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

		fmt.Println(sess.Get(sess.ID()))

		userById, err := r.UserStore.FindUserById(sess.Get(sess.ID()).(string))
		if err != nil {
			return err
		}

		return ctx.Status(http.StatusOK).JSON(userById)
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
