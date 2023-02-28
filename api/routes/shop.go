package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/api/middlewares"
	"planigo/pkg/shop"
)

func ShopRoutes(app fiber.Router, handler *shop.Handler) {
	router := app.Group("/shops")

	router.Get("/", handler.GetShops())
	router.Get("/:shopId", handler.GetShopById())
	router.Get("/owner/:ownerId", handler.GetShopsByUserId())

	router.Post("/",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin"}),
		handler.CreateShop(),
	)

	router.Patch("/:shopId",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.EditShop(),
	)

	router.Delete("/:shopId",
		middlewares.IsLoggedIn(handler.Session),
		middlewares.RequireRoles([]string{"admin", "owner"}),
		handler.DeleteShop(),
	)

	router.Get("/category/:categorySlug", handler.GetShopsByCategorySlug())
}
