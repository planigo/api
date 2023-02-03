package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/pkg/hour"
)

func HourRoutes(app fiber.Router, handler *hour.Handler) {
	hourRoutes := app.Group("/hours")

	hourRoutes.Get("/", handler.GetHours())
	hourRoutes.Post("/", handler.CreateHour())
	hourRoutes.Get("/:id", handler.GetHourById())
	hourRoutes.Delete("/:id", handler.DeleteHour())
	hourRoutes.Put("/:id", handler.UpdateHour())
}
