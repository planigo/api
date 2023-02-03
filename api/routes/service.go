package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/handlers"
)

func ServicesRoutes(app fiber.Router, handler *handlers.ServiceHandler) {
	router := app.Group("/services")

	router.Get("/", handler.GetServices())
	router.Get("/:serviceId", handler.GetServiceById())

	router.Post("/", handler.CreateService())

	router.Patch("/:serviceId", handler.EditService())
	router.Delete("/:serviceId", handler.DeleteService())
}
