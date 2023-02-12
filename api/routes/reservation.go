package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/pkg/reservation"
)

func ReservationRoutes(app fiber.Router, handler *reservation.Handler) {
	r := app.Group("/reservation")

	r.Get("slots/:shopId", handler.GetNextSlotsByDays())

	r.Post("/", handler.BookReservationByShopId())

	r.Post("/cancel/:id", handler.CancelReservation()) // can cancel only if the reservation is not started yet and it's owned by the user
}
