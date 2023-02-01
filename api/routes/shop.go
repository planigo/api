package routes

import (
	"planigo/pkg/shop"
	"github.com/gofiber/fiber/v2"
)

func ShopRoutes(app fiber.Router, handler *shop.ShopHandler) {
	router := app.Group("/shops")

	router.Get("/", handler.GetShops())
	router.Get("/:shopId", handler.GetShopById())

	router.Patch("/:shopId", handler.EditShop())

	router.Delete("/:shopId", handler.DeleteShop())

	// TODO POST /store createStore
	//router.Post("/", service.CreateStore)


}
