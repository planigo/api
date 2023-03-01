package tests

import (
	"github.com/stretchr/testify/assert"
	"planigo/utils"
	"testing"
	"time"
)

func TestIsNotOpeningHours(t *testing.T) {
	reservationAt := time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")
	isBetween := utils.IsReservationDuringOpeningHours(reservationAt, "10:00:00", "17:00:00", 60)
	assert.False(t, isBetween)
}

func TestIsDuringOpeningHours(t *testing.T) {
	reservationAt := time.Date(2021, 1, 1, 11, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")
	isBetween := utils.IsReservationDuringOpeningHours(reservationAt, "10:00:00", "17:00:00", 60)
	assert.True(t, isBetween)
}

func TestIsNotDuringOpeningHoursAfterClosingTime(t *testing.T) {
	reservationAt := time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")
	isBetween := utils.IsReservationDuringOpeningHours(reservationAt, "10:00:00", "17:00:00", 60)
	assert.False(t, isBetween)
}

func TestIsNotDuringOpeningHoursWithDurationExceedingClosingTime(t *testing.T) {
	reservationAt := time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")
	isBetween := utils.IsReservationDuringOpeningHours(reservationAt, "10:00:00", "17:00:00", 120)
	assert.False(t, isBetween)
}
