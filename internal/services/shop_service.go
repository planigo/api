package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"planigo/core/presenter"
	"planigo/internal/entities"
	"planigo/pkg/store"
)

type ShopHandler struct {
	*store.Store
}

func (h ShopHandler) GetShops() fiber.Handler {
	return func(c *fiber.Ctx) error {
		shops, err := h.ShopStore.FindShops()
		if err != nil {
			return presenter.Error(c, fiber.StatusNotFound, errors.New(presenter.RessourceNotFound))
		}

		return presenter.Response(c, fiber.StatusOK, shops)
	}
}

func (h ShopHandler) GetShopById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		shopId := c.Params("shopId")

		shop, err := h.ShopStore.FindShopById(shopId)
		if err != nil {
			return presenter.Error(c, fiber.StatusNotFound, errors.New(presenter.RessourceNotFound))
		}

		return presenter.Response(c, fiber.StatusOK, shop)
	}
}

func (h ShopHandler) GetShopsByUserId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ownerId := ctx.Params("ownerId")

		shops, err := h.ShopStore.FindShopsByUserId(ownerId)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(shops)
	}
}


func (h ShopHandler) CreateShop() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("userRole")

		if userRole != "admin" {
			return presenter.Error(c, fiber.StatusForbidden, errors.New(presenter.ActionNotAllowed))
		}
		newShop := new(entities.ShopRequest)
		if err := c.BodyParser(newShop); err != nil {
			return err
		}

		shop, err := h.ShopStore.AddShop(*newShop)
		if err != nil {
			return presenter.Error(c, fiber.StatusNotFound, errors.New(presenter.CannotAddShop))
		}

		return presenter.Response(c, fiber.StatusOK, shop)
	}
}

func (h ShopHandler) EditShop() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAllowedToUpdate := canUpdateShop(c, h)

		if !isAllowedToUpdate {
			return presenter.Error(c, fiber.StatusForbidden, errors.New(presenter.ActionNotAllowed))

		}
		shopEdited := new(entities.ShopRequest)
		shopId := c.Params("shopId")

		if err := c.BodyParser(shopEdited); err != nil {
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

		return presenter.Response(c, fiber.StatusOK, shop)
	}
}

func (h ShopHandler) DeleteShop() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAllowedToUpdate := canUpdateShop(c, h)

		if !isAllowedToUpdate {
			return presenter.Error(c, fiber.StatusForbidden, errors.New(presenter.ActionNotAllowed))
		}

		shopId := c.Params("shopId")

		code, err := h.ShopStore.RemoveShop(shopId)
		if err != nil {
			log.Fatal(code, err)
		}

		return presenter.Response(c, fiber.StatusOK, "")
	}
}

func (h ShopHandler) GetShopsByCategorySlug() fiber.Handler {
	return func(c *fiber.Ctx) error {
		categorySlug := c.Params("categorySlug")

		shops, err := h.ShopStore.FindShopsByCategorySlug(categorySlug)
		if err != nil {
			log.Fatal(err)
		}

		return presenter.Response(c, fiber.StatusOK, shops)
	}
}

func canUpdateShop(c *fiber.Ctx, h ShopHandler) bool {
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
