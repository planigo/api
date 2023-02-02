package shop

import (
	"log"
	"planigo/config/store"
	"planigo/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type ShopHandler struct {
	*store.Store
}

func NewHandler(store *store.Store) *ShopHandler {
	return &ShopHandler{store}
}

func (sh ShopHandler) GetShops() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shops, err := sh.ShopStore.FindShops()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shops)
	}
}

func (sh ShopHandler) GetShopById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shopId := ctx.Params("shopId")

		shop, err := sh.ShopStore.FindShopById(shopId)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shop)
	}
}

func (sh ShopHandler) CreateShop() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newShop := new(entities.ShopRequest)
		if err := ctx.BodyParser(newShop); err != nil {
			return err
		}

		shop, err := sh.ShopStore.AddShop(*newShop)

		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shop)
	}
}

func (sh ShopHandler) EditShop() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shopEdited := new(entities.ShopRequest)
		shopId := ctx.Params("shopId")

		if err := ctx.BodyParser(shopEdited); err != nil {
			return err
		}

		shopId, err := sh.ShopStore.UpdateShop(shopId, *shopEdited)
		if err != nil {
			log.Fatal(err)
		}

		shop, err := sh.ShopStore.FindShopById(shopId)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shop)
	}
}

func (sh ShopHandler) DeleteShop() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shopId := ctx.Params("shopId")

		code, err := sh.ShopStore.RemoveShop(shopId)
		if err != nil {
			log.Fatal(code, err)
		}

		return ctx.SendStatus(code)
	}
}
