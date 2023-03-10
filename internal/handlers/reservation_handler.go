package handlers

import (
	"planigo/core/enums"
	"planigo/core/middlewares"
	"planigo/internal/services"

	"github.com/gofiber/fiber/v2"
)

func ReservationRoutes(app fiber.Router, handler *services.ReservationHandler) {
	r := app.Group("/reservation")

	r.Get("slots/shop/:shopId",
		middlewares.IsLoggedIn,
		middlewares.RequireRoles([]string{enums.Admin}),
		handler.GetSlotsBookedByShop(),
	)

	r.Get(
		"slots/:shopId",
		handler.GetNextSlotsByDays(),
	)

	r.Get(
		"slots/users/:userId",
		handler.GetSlotsBookedByUser(),
	)

	r.Post(
		"/",
		handler.BookReservationByShopId(),
	)

	// can cancel only if the reservation is not started yet and it's owned by the user
	r.Get(
		"/cancel/:id",
		middlewares.IsLoggedIn,
		handler.CancelReservation(),
	)
}
