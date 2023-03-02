package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/core/middlewares"
	"planigo/internal/services"
)

func HourRoutes(app fiber.Router, handler *services.HourHandler) {
	hourRoutes := app.Group("/hours")

	hourRoutes.Get("/", handler.GetHours())
	hourRoutes.Get("/shop/:shopId", handler.GetHoursByShopId())
	hourRoutes.Post(
		"/",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.CreateHour(),
	)
	hourRoutes.Get("/:id", handler.GetHourById())
	hourRoutes.Delete(
		"/:id",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.DeleteHour(),
	)
	hourRoutes.Put(
		"/:id",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.UpdateHour(),
	)
}
