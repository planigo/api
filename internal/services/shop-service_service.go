package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
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
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
		}

		return ctx.JSON(services)
	}
}

func (sh ServiceHandler) GetServicesByShopId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		shopId := ctx.Params("shopId")
		services, err := sh.ServiceStore.FindServicesByShopId(shopId)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
		}

		return ctx.JSON(services)
	}
}

func (sh ServiceHandler) GetServiceById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		serviceId := ctx.Params("serviceId")

		service, err := sh.ServiceStore.FindServiceById(serviceId)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
		}

		return ctx.JSON(service)
	}
}

func (sh ServiceHandler) CreateService() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userRole := ctx.Locals("userRole")
		userId := ctx.Locals("userId")

		newService := new(entities.Service)
		shop, err := sh.ShopStore.FindShopById(newService.ShopID)

		if err != nil || (shop.OwnerID != userId && userRole != "admin") {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"statusCode": http.StatusUnauthorized,
				"message":    "You are not authorized to perform this action",
			})
		}

		if err := ctx.BodyParser(newService); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
		}

		serviceId, err := sh.ServiceStore.AddService(*newService)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
		}

		service, _ := sh.ServiceStore.FindServiceById(serviceId)

		return ctx.JSON(service)
	}
}

func (sh ServiceHandler) EditService() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isAllowedToUpdate := canUpdateService(ctx, sh)

		if !isAllowedToUpdate {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"statusCode": http.StatusUnauthorized,
				"message":    "You are not authorized to perform this action",
			})
		}

		serviceEdited := new(entities.Service)
		serviceId := ctx.Params("serviceId")

		if err := ctx.BodyParser(serviceEdited); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
		}

		serviceId, err := sh.ServiceStore.UpdateService(serviceId, *serviceEdited)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"statusCode": http.StatusInternalServerError,
				"message":    "User does not found",
			})
		}

		service, err := sh.ServiceStore.FindServiceById(serviceId)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"statusCode": http.StatusInternalServerError,
				"message":    fmt.Sprintf("Service %s not found", serviceId),
			})
		}

		return ctx.JSON(service)
	}
}

func (sh ServiceHandler) DeleteService() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isAllowedToUpdate := canUpdateService(ctx, sh)

		if !isAllowedToUpdate {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"statusCode": http.StatusUnauthorized,
				"message":    "You are not authorized to perform this action",
			})
		}
		serviceId := ctx.Params("serviceId")

		code, err := sh.ServiceStore.RemoveService(serviceId)
		if err != nil {
			log.Fatal(code, err)
		}

		return ctx.SendStatus(code)
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
