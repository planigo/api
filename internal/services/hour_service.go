package services

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"planigo/core/presenter"
	"planigo/internal/entities"
	"planigo/pkg/store"
)

type HourHandler struct {
	*store.Store
}

func (h HourHandler) GetHours() fiber.Handler {
	return func(c *fiber.Ctx) error {
		hours, err := h.HourStore.GetHours()
		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(c, fiber.StatusOK, hours)
	}
}

func (h HourHandler) GetHoursByShopId() fiber.Handler {
	return func(c *fiber.Ctx) error {
		shopId := c.Params("shopId")
		hours, err := h.FindHoursByShopId(shopId)
		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)

		}

		return presenter.Response(c, fiber.StatusOK, hours)

	}
}

func (h HourHandler) CreateHour() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("userRole")
		userId := c.Locals("userId")

		hour := parseHourBody(c)

		shop, err := h.ShopStore.FindShopById(hour.ShopID)

		if err != nil || (shop.OwnerID != userId && userRole != "admin") {
			return presenter.Error(c, fiber.StatusUnauthorized, errors.New(presenter.RessourceNotAuthorized))
		}

		createdHour, err := h.HourStore.CreateHour(*hour)
		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(c, fiber.StatusOK, createdHour)
	}
}

func (h HourHandler) GetHourById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		hours, err := h.HourStore.GetHourById(id)
		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(c, fiber.StatusOK, hours)
	}
}

func (h HourHandler) DeleteHour() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("DeleteHour")
		isAllowedToUpdate := canUpdateHour(c, h)

		if !isAllowedToUpdate {
			return presenter.Error(c, fiber.StatusUnauthorized, errors.New(presenter.RessourceNotAuthorized))
		}

		id := c.Params("id")
		err := h.HourStore.DeleteHour(id)
		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)

		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func (h HourHandler) UpdateHour() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAllowedToUpdate := canUpdateHour(c, h)

		if !isAllowedToUpdate {
			return presenter.Error(c, fiber.StatusUnauthorized, errors.New(presenter.RessourceNotAuthorized))
		}
		id := c.Params("id")
		updatedHour := parseHourBody(c)

		hour, err := h.HourStore.UpdateHour(id, *updatedHour)

		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(c, fiber.StatusOK, hour)
	}
}

func parseHourBody(c *fiber.Ctx) *entities.Hour {
	hour := new(entities.Hour)
	if err := c.BodyParser(hour); err != nil {
		log.Fatal(err)
	}

	return hour
}

func canUpdateHour(c *fiber.Ctx, h HourHandler) bool {
	userRole := c.Locals("userRole")
	userId := c.Locals("userId")
	id := c.Params("id")

	if userRole == "admin" {
		return true
	}

	hour, err := h.HourStore.GetHourById(id)
	if err != nil {
		return false
	}

	shop, err := h.ShopStore.FindShopById(hour.ShopID)

	if err != nil || (shop.OwnerID != userId) {
		return false
	}

	return true
}
