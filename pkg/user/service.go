package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"planigo/config/store"
	"planigo/pkg/entities"
)

type Handler struct {
	*store.Store
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{store}
}

func (r Handler) FindUsers() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := r.UserStore.FindUsers()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(users)

		return ctx.JSON(users)
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

		return ctx.JSON(userPayload)
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
