package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
	"planigo/config/mail"
	"planigo/config/store"
	"planigo/pkg/entities"
	user2 "planigo/pkg/user"
)

type Handler struct {
	*store.Store
	*mail.Mailer
	Session *session.Store
}

func New(store *store.Store, mailer *mail.Mailer, session *session.Store) *Handler {
	return &Handler{store, mailer, session}
}

func (r Handler) Login() fiber.Handler {
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

		if isSamePassword := user2.CheckPasswordHash(user.Password, findedUser.Password); !isSamePassword {
			return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Wrong email or password!",
			})
		}

		sess, err := r.Session.Get(ctx)
		if err != nil {
			panic(err)
		}

		sid := sess.ID()
		sess.Set(sid, findedUser.Id)
		sess.Set("role", findedUser.Role)

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
		fmt.Println(sess.Get("role"))

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

func (r Handler) ValidateEmail() fiber.Handler {
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
