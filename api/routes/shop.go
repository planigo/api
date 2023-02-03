package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/handlers"
)

func ShopRoutes(app fiber.Router, handler *handlers.ShopHandler) {
	router := app.Group("/shops")

	router.Get("/", handler.GetShops())
	router.Get("/:shopId", handler.GetShopById())

	router.Post("/", handler.CreateShop())

	router.Patch("/:shopId", handler.EditShop())

	router.Delete("/:shopId", handler.DeleteShop())
}
