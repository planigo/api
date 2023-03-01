package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/core/middlewares"
	"planigo/internal/services"
)

func ShopRoutes(app fiber.Router, handler *services.ShopHandler) {
	router := app.Group("/shops")

	router.Get("/", handler.GetShops())
	router.Get("/:shopId", handler.GetShopById())
	router.Get("/owner/:ownerId", handler.GetShopsByUserId())


	router.Post("/",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{"admin"}),
		handler.CreateShop(),
	)

	router.Patch("/:shopId",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.EditShop(),
	)

	router.Delete("/:shopId",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.DeleteShop(),
	)

	router.Get("/category/:categorySlug", handler.GetShopsByCategorySlug())
}
