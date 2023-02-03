package tests

import (
	"fmt"
	"planigo/common"
	"planigo/utils"
	"testing"
	"time"
)

var dateNow = time.Now()
var tomorrow = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 11, 0, 0, 0, time.UTC).AddDate(0, 0, 1)

func TestComputeEmptySlots(t *testing.T) {

	slots := utils.ComputeEmptySlots("10:00:00", "11:00:00")

	exceptedSlot := []common.Slot{
		{
			ReservationId: "",
			IsAvailable:   true,
			Start:         "10:00:00",
			End:           "11:00:00",
			Duration:      60,
		},
	}

	if slots[0].Start != exceptedSlot[0].Start {
		t.Errorf("start: got %s, want %s", slots[0].Start, exceptedSlot[0].Start)
	}

	if slots[0].End != exceptedSlot[0].End {
		t.Errorf("end: got %s, want %s", slots[0].End, exceptedSlot[0].End)
	}
}

func TestNextXDays(t *testing.T) {

	dates := utils.NextXDays(2)

	expectedDates := []string{
		dateNow.Format("2006-01-02"),
		dateNow.AddDate(0, 0, 1).Format("2006-01-02"),
		dateNow.AddDate(0, 0, 2).Format("2006-01-02"),
	}

	if dates[0].Format("2006-01-02") != expectedDates[0] || dates[1].Format("2006-01-02") != expectedDates[1] || dates[2].Format("2006-01-02") != expectedDates[2] {
		t.Errorf("got %+v, want %+v", dates, expectedDates)
	}

	dates = utils.NextXDays(3)

	expectedDates = append(expectedDates, dateNow.AddDate(0, 0, 3).Format("2006-01-02"))

	if dates[3].Format("2006-01-02") != expectedDates[3] {
		t.Errorf("got %+v, want %+v", dates, expectedDates)
	}
}

func TestCreateEmptySlotsMapByShopHours(t *testing.T) {

	emptySlotsMap := utils.CreateEmptySlotsMapByShopHours("10:00:00", "11:00:00", 2)

	expectedEmptySlotsMap := []common.DaySlot{
		{
			Date: time.Now().Format("2006-01-02"),
			Slots: []common.Slot{
				{
					ReservationId: "",
					IsAvailable:   true,
					Start:         "10:00:00",
					End:           "11:00:00",
					Duration:      60,
				},
			},
		},
		{
			Date: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
			Slots: []common.Slot{
				{
					ReservationId: "",
					IsAvailable:   true,
					Start:         "10:00:00",
					End:           "11:00:00",
					Duration:      60,
				},
			},
		},
		{
			Date: time.Now().AddDate(0, 0, 2).Format("2006-01-02"),
			Slots: []common.Slot{
				{
					ReservationId: "",
					IsAvailable:   true,
					Start:         "10:00:00",
					End:           "11:00:00",
					Duration:      60,
				},
			},
		},
	}

	if len(emptySlotsMap) != 3 {
		t.Errorf("len: got %d, want %d", len(emptySlotsMap), 3)
	}

	if emptySlotsMap[0].Date != expectedEmptySlotsMap[0].Date {
		t.Errorf("date 0: got %s, want %s", emptySlotsMap[0].Date, expectedEmptySlotsMap[0].Date)
	}

	if emptySlotsMap[1].Date != expectedEmptySlotsMap[1].Date {
		t.Errorf("date 1: got %s, want %s", emptySlotsMap[1].Date, expectedEmptySlotsMap[1].Date)
	}

	if emptySlotsMap[2].Date != expectedEmptySlotsMap[2].Date {
		t.Errorf("date 2: got %s, want %s", emptySlotsMap[2].Date, expectedEmptySlotsMap[2].Date)
	}
}

func TestMakeReservationMap(t *testing.T) {

	reservations := []common.DetailedReservation{
		{
			ReservationId: "1",
			ServiceId:     "1",
			ServiceName:   "coiffure homme",
			Start:         tomorrow.Format("2006-01-02 15:04:05"),
		},
	}

	reservationMap := utils.MakeReservationMap(reservations)

	key := fmt.Sprintf("%s-%s", tomorrow.Format("2006-01-02"), tomorrow.Format("15:04:05"))

	expectedReservationsMap := map[string]common.DetailedReservation{
		key: reservations[0],
	}

	if reservationMap[key].ReservationId != expectedReservationsMap[key].ReservationId {
		t.Errorf("got %s, want %s", reservationMap[key].ReservationId, expectedReservationsMap[key].ReservationId)
	}
}

func TestFillEmptySlotsWithReservationByDate(t *testing.T) {
	emptySlots := utils.CreateEmptySlotsMapByShopHours("11:00:00", "14:00:00", 2)

	daySlots := utils.FillEmptySlotsWithReservationByDate(emptySlots, []common.DetailedReservation{
		{
			ReservationId: "1",
			ServiceId:     "1",
			ServiceName:   "coiffure homme",
			Start:         tomorrow.Format("2006-01-02 15:04:05"),
		},
		{
			ReservationId: "1",
			ServiceId:     "1",
			ServiceName:   "coiffure homme",
			Start:         tomorrow.AddDate(0, 0, 1).Format("2006-01-02 15:04:05"),
		},
	})

	if daySlots[1].Slots[0].IsAvailable {
		t.Errorf("got %t, want %t", daySlots[1].Slots[0].IsAvailable, false)
	}

	if !daySlots[1].Slots[1].IsAvailable {
		t.Errorf("got %t, want %t", daySlots[1].Slots[0].IsAvailable, true)
	}

	if daySlots[2].Slots[0].IsAvailable {
		t.Errorf("got %t, want %t", daySlots[1].Slots[0].IsAvailable, false)
	}

	if !daySlots[2].Slots[1].IsAvailable {
		t.Errorf("got %t, want %t", daySlots[1].Slots[0].IsAvailable, true)
	}
}
