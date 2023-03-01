package reservation

import (
	"fmt"
	"log"
	"net/http"
	"planigo/common"
	"planigo/config/mail"
	"planigo/config/store"
	"planigo/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"

)

type Handler struct {
	*store.Store
	Session *session.Store
	*mail.Mailer
}

func New(store *store.Store, session *session.Store, mailer *mail.Mailer) *Handler {
	return &Handler{store, session, mailer}
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

func (h Handler) GetSlotsBookedByUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId")
		fmt.Println(userId)

		bookedReservation, err := h.ReservationStore.GetSlotsBookedByUserId(userId)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.ErrInternalServerError,
				"message":    err.Error(),
			})
		}

		return ctx.Status(http.StatusOK).JSON(bookedReservation)
	}
}

func (h Handler) BookReservationByShopId() fiber.Handler {
	body := struct {
		ServiceId string `json:"serviceId"`
		ShopId    string `json:"shopId"`
		Start     string `json:"start"`
		UserId    string `json:"userId"`
	}{}
	return func(ctx *fiber.Ctx) error {
		if err := ctx.BodyParser(&body); err != nil {
			log.Println(err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"message":    err.Error(),
			})
		}

		reservationAt, _ := time.Parse("2006-01-02 15:04:05", body.Start)
		shopHourForWantedDay, err := h.HourStore.GetHourByShopIdAndDay(body.ShopId, utils.GetDayNumberWithSundayAsLast(int(reservationAt.Weekday())))
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"statusCode": fiber.StatusNotFound,
				"message":    err.Error(),
			})
		}

		if !utils.IsReservationDuringOpeningHours(body.Start, shopHourForWantedDay.Start, shopHourForWantedDay.End, 60) {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"statusCode": fiber.StatusForbidden,
				"message":    fmt.Sprintf("The shop is closed at %s", reservationAt.Format("15:04")),
			})
		}

		reservation, err := h.ReservationStore.BookReservation(body.ServiceId, body.ShopId, body.Start, body.UserId)
		if err != nil {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"statusCode": fiber.StatusConflict,
				"message":    err.Error(),
			})
		}

		user, err := h.UserStore.FindUserById(reservation.UserId)
		if err != nil {
			log.Println(err.Error())
		}

		reservationBookedEmail := mail.Content{
			To:      user.Email,
			Subject: "Reservation booked",
			Body:    getEmailContentForBookedReservation(reservation),
		}
		if err := h.Mailer.Send(reservationBookedEmail); err != nil {
			log.Println(err.Error())
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

		user, err := h.UserStore.FindUserById(reservation.UserId)
		if err != nil {
			log.Println(err.Error())
		}

		reservationCanceledEmail := mail.Content{
			To:      user.Email,
			Subject: "Reservation canceled",
			Body:    getEmailContentForCancelation(reservation),
		}
		if err := h.Mailer.Send(reservationCanceledEmail); err != nil {
			log.Println(err.Error())
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

func getEmailContentForBookedReservation(reservation common.DetailedReservation) string {
	return fmt.Sprintf("Your reservation for %s at %s has been booked", reservation.ServiceName, reservation.Start)
}

func getEmailContentForCancelation(reservation common.DetailedReservation) string {
	return fmt.Sprintf("Your reservation for %s at %s has been canceled", reservation.ServiceName, reservation.Start)
}
