package tests

import (
	"planigo/utils"
	"testing"
	"time"
)

func TestIsNotOpeningHours(t *testing.T) {
	reservationAt := time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")
	isBetween := utils.IsReservationDuringOpeningHours(reservationAt, "10:00:00", "17:00:00", 60)
	if isBetween != false {
		t.Errorf("got %+v, want %+v", isBetween, false)
	}
}

func TestIsDuringOpeningHours(t *testing.T) {
	reservationAt := time.Date(2021, 1, 1, 11, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")
	isBetween := utils.IsReservationDuringOpeningHours(reservationAt, "10:00:00", "17:00:00", 60)
	if isBetween != true {
		t.Errorf("got %+v, want %+v", isBetween, true)
	}
}

func TestIsNotDuringOpeningHoursAfterClosingTime(t *testing.T) {
	reservationAt := time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")
	isBetween := utils.IsReservationDuringOpeningHours(reservationAt, "10:00:00", "17:00:00", 60)
	if isBetween != false {
		t.Errorf("got %+v, want %+v", isBetween, false)
	}
}

func TestIsNotDuringOpeningHoursWithDurationExceedingClosingTime(t *testing.T) {
	reservationAt := time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")
	isBetween := utils.IsReservationDuringOpeningHours(reservationAt, "10:00:00", "17:00:00", 120)
	if isBetween != false {
		t.Errorf("got %+v, want %+v", isBetween, false)
	}
}
