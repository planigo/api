package routes

import (
	"planigo/pkg/hour"

	"github.com/gofiber/fiber/v2"
)

func HourRoutes(app fiber.Router, handler *hour.Handler) {
	hourRoutes := app.Group("/hours")

	hourRoutes.Get("/", handler.GetHours())
	hourRoutes.Get("/:shopId", handler.GetHoursByShopId())
	hourRoutes.Post("/", handler.CreateHour())
	hourRoutes.Get("/:id", handler.GetHourById())
	hourRoutes.Delete("/:id", handler.DeleteHour())
	hourRoutes.Put("/:id", handler.UpdateHour())
}
