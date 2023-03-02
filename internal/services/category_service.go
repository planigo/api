package services

import (
	"github.com/gofiber/fiber/v2"
	"planigo/core/presenter"
	"planigo/pkg/store"
)

type CategoryHandler struct {
	*store.Store
}

func (h CategoryHandler) GetCategories() fiber.Handler {
	return func(c *fiber.Ctx) error {
		categories, err := h.CategoryStore.GetCategories()
		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(c, fiber.StatusOK, categories)
	}
}
