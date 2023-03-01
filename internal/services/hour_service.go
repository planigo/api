package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
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
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return c.JSON(hours)
	}
}

func (h HourHandler) GetHoursByShopId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shopId := ctx.Params("shopId")
		hours, err := h.FindHoursByShopId(shopId)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return ctx.JSON(hours)
	}
}

func (h HourHandler) CreateHour() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("userRole")
		userId := c.Locals("userId")

		hour := parseHourBody(c)

		shop, err := h.ShopStore.FindShopById(hour.ShopID)

		if err != nil || (shop.OwnerID != userId && userRole != "admin") {
			return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You are not authorized to create this resource",
			})
		}

		createdHour, err := h.HourStore.CreateHour(*hour)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return c.JSON(createdHour)
	}
}

func (h HourHandler) GetHourById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		hours, err := h.HourStore.GetHourById(id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return c.JSON(hours)
	}
}

func (h HourHandler) DeleteHour() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("DeleteHour")
		isAllowedToUpdate := canUpdateHour(c, h)

		if !isAllowedToUpdate {
			return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You are not authorized to delete this resource",
			})
		}

		id := c.Params("id")
		err := h.HourStore.DeleteHour(id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (h HourHandler) UpdateHour() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAllowedToUpdate := canUpdateHour(c, h)

		if !isAllowedToUpdate {
			return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You are not authorized to update this resource",
			})
		}
		id := c.Params("id")
		updatedHour := parseHourBody(c)

		hour, err := h.HourStore.UpdateHour(id, *updatedHour)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return c.JSON(hour)
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
