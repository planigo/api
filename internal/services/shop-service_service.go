package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"planigo/core/presenter"
	"planigo/internal/entities"
	"planigo/pkg/store"
)

type ServiceHandler struct {
	*store.Store
}

func (sh ServiceHandler) GetServices() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		services, err := sh.ServiceStore.FindServices()
		if err != nil {
			return presenter.Error(ctx, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(ctx, fiber.StatusOK, services)
	}
}

func (sh ServiceHandler) GetServicesByShopId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shopId := ctx.Params("shopId")
		services, err := sh.ServiceStore.FindServicesByShopId(shopId)
		if err != nil {
			return presenter.Error(ctx, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(ctx, fiber.StatusOK, services)
	}
}

func (sh ServiceHandler) GetServiceById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		serviceId := ctx.Params("serviceId")

		service, err := sh.ServiceStore.FindServiceById(serviceId)

		if err != nil {
			return presenter.Error(ctx, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(ctx, fiber.StatusOK, service)
	}
}

func (sh ServiceHandler) CreateService() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userRole := ctx.Locals("userRole")
		userId := ctx.Locals("userId")

		newService := new(entities.Service)
		shop, err := sh.ShopStore.FindShopById(newService.ShopID)

		if err != nil || (shop.OwnerID != userId && userRole != "admin") {
			return presenter.Error(ctx, fiber.StatusInternalServerError, errors.New(presenter.ActionNotAllowed))

		}

		if err := ctx.BodyParser(newService); err != nil {
			return presenter.Error(ctx, fiber.StatusInternalServerError, err)
		}

		serviceId, err := sh.ServiceStore.AddService(*newService)

		if err != nil {
			return presenter.Error(ctx, fiber.StatusInternalServerError, err)
		}

		service, _ := sh.ServiceStore.FindServiceById(serviceId)

		return presenter.Response(ctx, fiber.StatusOK, service)
	}
}

func (sh ServiceHandler) EditService() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isAllowedToUpdate := canUpdateService(ctx, sh)

		if !isAllowedToUpdate {
			return presenter.Error(ctx, fiber.StatusForbidden, errors.New(presenter.ActionNotAllowed))
		}

		serviceEdited := new(entities.Service)
		serviceId := ctx.Params("serviceId")

		if err := ctx.BodyParser(serviceEdited); err != nil {
			return presenter.Error(ctx, fiber.StatusInternalServerError, err)
		}

		serviceId, err := sh.ServiceStore.UpdateService(serviceId, *serviceEdited)
		if err != nil {
			return presenter.Error(ctx, fiber.StatusNotFound, err)
		}

		service, err := sh.ServiceStore.FindServiceById(serviceId)
		if err != nil {
			return presenter.Error(ctx, fiber.StatusNotFound, errors.New(presenter.RessourceNotFound))
		}

		return presenter.Response(ctx, fiber.StatusOK, service)
	}
}

func (sh ServiceHandler) DeleteService() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isAllowedToUpdate := canUpdateService(ctx, sh)

		if !isAllowedToUpdate {
			return presenter.Error(ctx, fiber.StatusForbidden, errors.New(presenter.ActionNotAllowed))
		}
		serviceId := ctx.Params("serviceId")

		_, err := sh.ServiceStore.RemoveService(serviceId)
		if err != nil {
			return presenter.Error(ctx, fiber.StatusForbidden, errors.New(presenter.CannotRemoveService))
		}

		return presenter.Response(ctx, fiber.StatusOK, "")
	}

}

func canUpdateService(c *fiber.Ctx, h ServiceHandler) bool {
	userRole := c.Locals("userRole")
	userId := c.Locals("userId")
	id := c.Params("id")

	if userRole == "admin" {
		return true
	}

	service, err := h.ServiceStore.FindServiceById(id)

	if err != nil {
		return false
	}

	shop, err := h.ShopStore.FindShopById(service.ShopID)

	if err != nil || shop.OwnerID != userId {
		return false
	}

	return true
}
