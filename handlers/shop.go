package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"planigo/config/store"
	"planigo/models"
)

type ShopHandler struct {
	*store.Store
	Session *session.Store
}

func (h ShopHandler) GetShops() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shops, err := h.ShopStore.FindShops()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shops)
	}
}

func (h ShopHandler) GetShopById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shopId := ctx.Params("shopId")

		shop, err := h.ShopStore.FindShopById(shopId)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shop)
	}
}

func (h ShopHandler) CreateShop() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newShop := new(models.ShopRequest)
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

func (h ShopHandler) EditShop() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shopEdited := new(models.ShopRequest)
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

func (h ShopHandler) DeleteShop() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shopId := ctx.Params("shopId")

		code, err := h.ShopStore.RemoveShop(shopId)
		if err != nil {
			log.Fatal(code, err)
		}

		return ctx.SendStatus(code)
	}
}
