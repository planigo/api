package utils

import (
	"planigo/common"
	"time"
)

func CreateEmptySlotsMapByShopHours(startHour string, endHour string) []common.DaySlot {
	var emptySlotsMap []common.DaySlot

	for _, day := range NextSevenDays() {
		var slot common.DaySlot
		slot = common.DaySlot{
			Date:  day.Format("2006-01-02"),
			Slots: ComputeEmptySlots(startHour, endHour),
		}
		emptySlotsMap = append(emptySlotsMap, slot)
	}

	return emptySlotsMap
}

func NextSevenDays() []time.Time {
	var dates []time.Time
	for i := 0; i < 7; i++ {
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

func FillEmptySlotsWithRevervationByDate(
	emptySlotsMap []common.DaySlot,
	reservations []common.DetailledReservation,
) []common.DaySlot {

	reservationMap := makeReservationMap(reservations)

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

func makeReservationMap(
	reservations []common.DetailledReservation,
) map[string]common.DetailledReservation {
	reservationMap := make(map[string]common.DetailledReservation)

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
