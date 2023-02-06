package category

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"planigo/config/store"
)

type Handler struct {
	*store.Store
}

func (h Handler) GetCategories() fiber.Handler {
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
