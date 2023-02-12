package reservation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"net/http"
	"planigo/config/store"
	"planigo/utils"
	"strconv"
)

type Handler struct {
	*store.Store
	Session *session.Store
}

func New(store *store.Store, session *session.Store) *Handler {
	return &Handler{store, session}
}

func (h Handler) GetNextSlotsByDays() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		shopId := ctx.Params("shopId")
		nbOfDays, _ := strconv.Atoi(ctx.Query("until", "7"))

		bookedReservations, err := h.ReservationStore.GetReservationsByShopId(shopId)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"message":    err.Error(),
			})
		}

		shopHoursByWeekDay, err := h.HourStore.GetHourByShopId(shopId)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.ErrNotFound,
				"message":    err.Error(),
			})
		}

		emptySlots := utils.CreateEmptySlotsWithShopHours(shopHoursByWeekDay, nbOfDays)

		return ctx.
			Status(http.StatusOK).
			JSON(utils.FillEmptySlotsWithReservationByDate(emptySlots, bookedReservations))
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

func (h Handler) CancelReservation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reservationId := c.Params("id")
		userId := c.Locals("userId")

		reservation, err := h.ReservationStore.GetReservationById(reservationId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"statusCode": fiber.StatusNotFound,
				"message":    err.Error(),
			})
		}

		// only the user who booked the reservation can cancel it > move to isOwner() middleware ?
		if reservation.UserId != userId {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": fiber.StatusUnauthorized,
				"message":    "You are not allowed to cancel this reservation",
			})
		}

		if err := h.ReservationStore.CancelReservation(reservationId); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"message":    err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"statusCode": http.StatusOK,
			"message":    "Reservation canceled",
		})
	}
}
