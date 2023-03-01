package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"net/http"
	"planigo/config/store"
	"planigo/internal/entities"
)

type ServiceHandler struct {
	*store.Store
	Session *session.Store
}

func NewServiceHandler(store *store.Store, session *session.Store) *ServiceHandler {
	return &ServiceHandler{store, session}
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
		newService := new(entities.Service)
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
		serviceId := ctx.Params("serviceId")

		code, err := sh.ServiceStore.RemoveService(serviceId)
		if err != nil {
			log.Fatal(code, err)
		}

		return ctx.SendStatus(code)
	}

}
