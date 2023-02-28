package shop

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"planigo/config/store"
	"planigo/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	*store.Store
	Session *session.Store
}

func New(store *store.Store, session *session.Store) *Handler {
	return &Handler{store, session}
}

func (h Handler) GetShops() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shops, err := h.ShopStore.FindShops()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shops)
	}
}

func (h Handler) GetShopById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shopId := ctx.Params("shopId")

		shop, err := h.ShopStore.FindShopById(shopId)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shop)
	}
}

func (h Handler) GetShopsByUserId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ownerId := ctx.Params("ownerId")

		shops, err := h.ShopStore.FindShopsByUserId(ownerId)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shops)
	}
}

func (h Handler) CreateShop() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userRole := ctx.Locals("userRole")

		if userRole != "admin" {
			return ctx.Status(401).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You are not authorized to create this resource",
			})
		}
		newShop := new(entities.ShopRequest)
		if err := ctx.BodyParser(newShop); err != nil {
			return err
		}

		shop, err := h.ShopStore.AddShop(*newShop)

		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shop)
	}
}

func (h Handler) EditShop() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isAllowedToUpdate := canUpdateShop(ctx, h)

		if !isAllowedToUpdate {
			return ctx.Status(401).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You are not authorized to update this resource",
			})
		}
		shopEdited := new(entities.ShopRequest)
		shopId := ctx.Params("shopId")

		if err := ctx.BodyParser(shopEdited); err != nil {
			return err
		}

		shopId, err := h.ShopStore.UpdateShop(shopId, *shopEdited)
		if err != nil {
			log.Fatal(err)
		}

		shop, err := h.ShopStore.FindShopById(shopId)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shop)
	}
}

func (h Handler) DeleteShop() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isAllowedToUpdate := canUpdateShop(ctx, h)

		if !isAllowedToUpdate {
			return ctx.Status(401).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You are not authorized to update this resource",
			})
		}

		shopId := ctx.Params("shopId")

		code, err := h.ShopStore.RemoveShop(shopId)
		if err != nil {
			log.Fatal(code, err)
		}

		return ctx.SendStatus(code)
	}
}

func (h Handler) GetShopsByCategorySlug() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		categorySlug := ctx.Params("categorySlug")

		shops, err := h.ShopStore.FindShopsByCategorySlug(categorySlug)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shops)
	}
}

func canUpdateShop(c *fiber.Ctx, h Handler) bool {
	shopId := c.Params("shopId")
	userId := c.Locals("userId")
	userRole := c.Locals("userRole")

	if userRole == "admin" {
		return true
	}

	shop, err := h.ShopStore.FindShopById(shopId)
	if err != nil || (shop.OwnerID != userId) {
		return false
	}

	return true
}
