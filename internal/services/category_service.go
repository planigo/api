package services

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"planigo/pkg/store"
)

type CategoryHandler struct {
	*store.Store
}

func (h CategoryHandler) GetCategories() fiber.Handler {
	return func(c *fiber.Ctx) error {
		categories, err := h.CategoryStore.GetCategories()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return c.JSON(categories)
	}
}
