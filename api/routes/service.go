package routes

import (
	"planigo/api/middlewares"
	"planigo/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func ServicesRoutes(app fiber.Router, handler *service.ServiceHandler) {
	router := app.Group("/services")

	router.Get("/", handler.GetServices())
	router.Get("/shop/:shopId", handler.GetServicesByShopId())
	router.Get("/:serviceId", handler.GetServiceById())

	router.Post(
		"/",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.CreateService(),
	)

	router.Patch("/:serviceId",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.EditService(),
	)
	router.Delete("/:serviceId",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.DeleteService(),
	)
}
