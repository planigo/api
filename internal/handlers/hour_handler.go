package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/internal/services"
)

func HourRoutes(app fiber.Router, handler *services.HourHandler) {
	hourRoutes := app.Group("/hours")

	hourRoutes.Get("/", handler.GetHours())
	hourRoutes.Get("/:shopId", handler.GetHoursByShopId())
	hourRoutes.Post("/", handler.CreateHour())
	hourRoutes.Get("/:id", handler.GetHourById())
	hourRoutes.Delete("/:id", handler.DeleteHour())
	hourRoutes.Put("/:id", handler.UpdateHour())
}
