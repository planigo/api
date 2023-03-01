package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/internal/services"
)

func ServicesRoutes(app fiber.Router, handler *services.ServiceHandler) {
	router := app.Group("/services")

	router.Get("/", handler.GetServices())
	router.Get("/shop/:shopId", handler.GetServicesByShopId())
	router.Get("/:serviceId", handler.GetServiceById())

	router.Post("/", handler.CreateService())

	router.Patch("/:serviceId", handler.EditService())
	router.Delete("/:serviceId", handler.DeleteService())
}
