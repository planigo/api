package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/internal/services"
)

func ShopRoutes(app fiber.Router, handler *services.ShopHandler) {
	router := app.Group("/shops")

	router.Get("/", handler.GetShops())
	router.Get("/:shopId", handler.GetShopById())

	router.Post("/", handler.CreateShop())

	router.Patch("/:shopId", handler.EditShop())

	router.Delete("/:shopId", handler.DeleteShop())

	router.Get("/category/:categorySlug", handler.GetShopsByCategorySlug())
}
