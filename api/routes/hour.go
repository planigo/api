package routes

import (
	"planigo/api/middlewares"
	"planigo/pkg/hour"

	"github.com/gofiber/fiber/v2"
)

func HourRoutes(app fiber.Router, handler *hour.Handler) {
	hourRoutes := app.Group("/hours")

	hourRoutes.Get("/", handler.GetHours())
	hourRoutes.Get("/:shopId", handler.GetHoursByShopId())
	hourRoutes.Post(
		"/",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.CreateHour(),
	)
	hourRoutes.Get("/:id", handler.GetHourById())
	hourRoutes.Delete(
		"/:id",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.DeleteHour(),
	)
	hourRoutes.Put(
		"/:id",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.UpdateHour(),
	)
}
