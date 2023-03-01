package handlers

import (
	"github.com/gofiber/fiber/v2"
	"planigo/core/middlewares"
	"planigo/internal/services"
)

func ReservationRoutes(app fiber.Router, handler *services.ReservationHandler) {
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
