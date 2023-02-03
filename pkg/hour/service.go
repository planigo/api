package hour

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"net/http"
	"planigo/config/store"
	"planigo/pkg/entities"
	"strconv"
)

type Handler struct {
	*store.Store
	Session *session.Store
}

func (h Handler) GetHours() fiber.Handler {
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

func (h Handler) CreateHour() fiber.Handler {
	return func(c *fiber.Ctx) error {
		hour := parseHourBody(c)

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

func (h Handler) GetHourById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

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

func (h Handler) DeleteHour() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		println(id)

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

func (h Handler) UpdateHour() fiber.Handler {
	return func(c *fiber.Ctx) error {
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
