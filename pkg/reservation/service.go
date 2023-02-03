package reservation

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"net/http"
	"planigo/config/store"
	"planigo/utils"
)

type Handler struct {
	*store.Store
	Session *session.Store
}

func New(store *store.Store, session *session.Store) *Handler {
	return &Handler{store, session}
}

func (h Handler) GetResevationByShopId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		shopId := ctx.Params("shopId")

		println(shopId)

		bookedReservations, err := h.ReservationStore.GetReservationsByShopId(shopId)
		if err != nil {
			fmt.Println(err.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"message":    err.Error(),
			})
		}

		//shopHoursByWeekDay, err = h.HourStore.GetHoursByShopId()
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		//		"statusCode": fiber.StatusInternalServerError,
		//		"message":    err.Error(),
		//	})
		//}

		emptySlots := utils.CreateEmptySlotsMapByShopHours("09:00:00", "18:00:00")

		return ctx.
			Status(http.StatusOK).
			JSON(&fiber.Map{
				"data": utils.FillEmptySlotsWithRevervationByDate(emptySlots, bookedReservations),
			})
	}
}

func (h Handler) BookReservationByShopId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		body := struct {
			ServiceId string `json:"serviceId"`
			ShopId    string `json:"shopId"`
			Start     string `json:"start"`
			UserId    string `json:"userId"`
		}{}
		if err := ctx.BodyParser(&body); err != nil {
			log.Println(err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"message":    err.Error(),
			})
		}

		reservation, err := h.ReservationStore.BookReservation(body.ServiceId, body.ShopId, body.Start, body.UserId)
		if err != nil {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"statusCode": fiber.StatusConflict,
				"message":    err.Error(),
			})
		}

		return ctx.Status(http.StatusCreated).JSON(reservation)
	}

}
