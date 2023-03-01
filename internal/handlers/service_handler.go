package handlers

import (
	"github.com/gofiber/fiber/v2"

	"planigo/core/middlewares"
	"planigo/internal/services"
)

func ServicesRoutes(app fiber.Router, handler *services.ServiceHandler) {
	router := app.Group("/services")

	router.Get("/", handler.GetServices())
	router.Get("/shop/:shopId", handler.GetServicesByShopId())
	router.Get("/:serviceId", handler.GetServiceById())

	router.Post(
		"/",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.CreateService(),
	)

	router.Patch("/:serviceId",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.EditService(),
	)
	router.Delete("/:serviceId",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.DeleteService(),
	)
}
