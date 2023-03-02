package services

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"planigo/common"
	"planigo/core/presenter"
	"planigo/pkg/mail"
	"planigo/pkg/store"
	"planigo/utils"
	"strconv"
	"time"
)

type ReservationHandler struct {
	*store.Store
	*mail.Mailer
}

func (h ReservationHandler) GetNextSlotsByDays() fiber.Handler {
	return func(c *fiber.Ctx) error {

		shopId := c.Params("shopId")
		nbOfDays, _ := strconv.Atoi(c.Query("until", "7"))

		bookedReservations, err := h.ReservationStore.GetReservationsByShopId(shopId)
		if err != nil {
			return presenter.Response(c, fiber.StatusInternalServerError, err)
		}

		shopHoursByWeekDay, err := h.HourStore.GetHourByShopId(shopId)
		if err != nil {
			return presenter.Response(c, fiber.StatusNotFound, errors.New(presenter.NoHourForThisDay))
		}

		emptySlots := utils.CreateEmptySlotsWithShopHours(shopHoursByWeekDay, nbOfDays)

		return presenter.Response(c, fiber.StatusOK, utils.FillEmptySlotsWithReservationByDate(emptySlots, bookedReservations))
	}
}

func (h ReservationHandler) GetSlotsBookedByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("userId")

		bookedReservation, err := h.ReservationStore.GetSlotsBookedByUserId(userId)
		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(c, fiber.StatusOK, bookedReservation)
	}
}

func (h ReservationHandler) GetSlotsBookedByShop() fiber.Handler {
	return func(c *fiber.Ctx) error {
		shopId := c.Params("shopId")

		reservations, err := h.ReservationStore.FindSlotsBookedFilteredShopId(shopId)
		if err != nil {
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(c, fiber.StatusOK, reservations)
	}
}

func (h ReservationHandler) BookReservationByShopId() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := struct {
			ServiceId string `json:"serviceId"`
			ShopId    string `json:"shopId"`
			Start     string `json:"start"`
			UserId    string `json:"userId"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return presenter.Error(c, fiber.StatusBadRequest, err)
		}

		reservationAt, _ := time.Parse("2006-01-02 15:04:05", body.Start)
		shopHourForWantedDay, err := h.HourStore.GetHourByShopIdAndDay(body.ShopId, utils.GetDayNumberWithSundayAsLast(int(reservationAt.Weekday())))
		if err != nil {
			return presenter.Error(c, fiber.StatusNotFound, errors.New(presenter.NoHourForThisDay))
		}

		if !utils.IsReservationDuringOpeningHours(body.Start, shopHourForWantedDay.Start, shopHourForWantedDay.End, 60) {
			return presenter.Error(c, fiber.StatusForbidden, errors.New(presenter.ReservationOutOfOpeningHours))
		}

		reservation, err := h.ReservationStore.BookReservation(body.ServiceId, body.ShopId, body.Start, body.UserId)
		if err != nil {
			return presenter.Error(c, fiber.StatusConflict, err)
		}

		user, err := h.UserStore.FindUserById(body.UserId)

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

		return presenter.Response(c, fiber.StatusCreated, reservation)
	}
}

func (h ReservationHandler) CancelReservation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reservationId := c.Params("id")
		userId := c.Locals("userId")

		reservation, err := h.ReservationStore.GetReservationById(reservationId)
		if err != nil {
			return presenter.Error(c, fiber.StatusNotFound, err)
		}

		// only the user who booked the reservation can cancel it > move to isOwner() middleware ?
		if reservation.UserId != userId {
			return presenter.Error(c, fiber.StatusUnauthorized, errors.New(presenter.NotAllowedToCancelReservation))
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
			return presenter.Error(c, fiber.StatusInternalServerError, err)
		}

		return presenter.Response(c, fiber.StatusOK, presenter.ReservationCanceled)
	}
}

func getEmailContentForBookedReservation(reservation common.DetailedReservation) string {
	return fmt.Sprintf("Your reservation for %s at %s has been booked", reservation.ServiceName, reservation.Start)
}

func getEmailContentForCancelation(reservation common.DetailedReservation) string {
	return fmt.Sprintf("Your reservation for %s at %s has been canceled", reservation.ServiceName, reservation.Start)
}
