package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/api/middlewares"
	"planigo/pkg/reservation"
)

func ReservationRoutes(app fiber.Router, handler *reservation.Handler) {
	r := app.Group("/reservation")

	r.Get(
		"slots/:shopId",
		handler.GetNextSlotsByDays(),
	)

	r.Post(
		"/",
		handler.BookReservationByShopId(),
	)

	// can cancel only if the reservation is not started yet and it's owned by the user
	r.Get(
		"/cancel/:id",
		middlewares.IsLoggedIn(handler.Session),
		handler.CancelReservation(),
	)
}
