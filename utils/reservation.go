package utils

import (
	"planigo/common"
	"time"
)

func CreateEmptySlotsMapByShopHours(startHour string, endHour string, days int) []common.DaySlot {
	var emptySlotsMap []common.DaySlot

	for _, day := range NextXDays(days) {
		var slot common.DaySlot
		slot = common.DaySlot{
			Date:  day.Format("2006-01-02"),
			Slots: ComputeEmptySlots(startHour, endHour),
		}
		emptySlotsMap = append(emptySlotsMap, slot)
	}

	return emptySlotsMap
}

func NextXDays(x int) []time.Time {
	var dates []time.Time
	for i := 0; i <= x; i++ {
		date := time.Now().AddDate(0, 0, i)
		dates = append(dates, date)
	}
	return dates
}

func ComputeEmptySlots(startHour string, endHour string) []common.Slot {
	var slots []common.Slot

	start, _ := time.Parse("15:04:05", startHour)
	end, _ := time.Parse("15:04:05", endHour)

	for i := start; i.Before(end); i = i.Add(time.Hour) {
		slot := common.Slot{
			ReservationId: "",
			IsAvailable:   true,
			Start:         i.Format("15:04:05"),
			End:           i.Add(time.Hour).Format("15:04:05"),
			Duration:      60,
		}
		slots = append(slots, slot)
	}

	return slots
}

func FillEmptySlotsWithReservationByDate(
	emptySlotsMap []common.DaySlot,
	reservations []common.DetailedReservation,
) []common.DaySlot {

	reservationMap := MakeReservationMap(reservations)

	for i := range emptySlotsMap {
		for j := range emptySlotsMap[i].Slots {
			key := getReservationKey(emptySlotsMap[i].Date, emptySlotsMap[i].Slots[j].Start)
			if reservation, ok := reservationMap[key]; ok {
				emptySlotsMap[i].Slots[j].ReservationId = reservation.ReservationId
				emptySlotsMap[i].Slots[j].IsAvailable = false
			}
		}
	}

	return emptySlotsMap
}

func MakeReservationMap(
	reservations []common.DetailedReservation,
) map[string]common.DetailedReservation {
	reservationMap := make(map[string]common.DetailedReservation)

	for _, reservation := range reservations {
		reservationDate, _ := time.Parse("2006-01-02 15:04:05", reservation.Start)
		key := getReservationKey(reservationDate.Format("2006-01-02"), reservationDate.Format("15:04:05"))
		reservationMap[key] = reservation
	}
	return reservationMap
}

func getReservationKey(date string, hour string) string {
	return date + "-" + hour
}
