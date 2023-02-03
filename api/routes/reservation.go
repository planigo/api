package routes

import (
	"github.com/gofiber/fiber/v2"
	"planigo/pkg/reservation"
)

func ReservationRoutes(app fiber.Router, handler *reservation.Handler) {
	r := app.Group("/reservation")

	r.Get("/:shopId", handler.GetResevationByShopId())

	r.Post("/", handler.BookReservationByShopId())

	//r.Get("/:id", handler.GetHourById())
	//r.Delete("/:id", handler.DeleteHour())
	//r.Put("/:id", handler.UpdateHour())
}
