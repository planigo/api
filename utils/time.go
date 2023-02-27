package utils

import (
	"fmt"
	"math"
	"time"
)

// IsReservationDuringOpeningHours check if the wanted hour is in the shop hours
// check also if the service duration is not over the shop closing hour
// can't book at 17:00 if the service duration is 120 minutes and the shop is closing at 17:00
func IsReservationDuringOpeningHours(reservationAt string, start string, end string, serviceDurationInMinutes int) bool {
	durationInHour := math.Round(float64(serviceDurationInMinutes / 60))
	reservationStart, err := time.Parse("2006-01-02 15:04:05", reservationAt)
	if err != nil {
		fmt.Println("reservationStart :", err.Error())
		return false
	}
	shopStartHour, err := time.Parse("15:04:00", start)
	if err != nil {
		fmt.Println("shopStartHour :", err.Error())
		return false
	}
	shopEndHour, err := time.Parse("15:04:00", end)
	if err != nil {
		fmt.Println("shopEndHour :", err.Error())
		return false
	}

	return reservationStart.Hour() >= shopStartHour.Hour() && reservationStart.Hour() < shopEndHour.Hour()-int(durationInHour)
}
