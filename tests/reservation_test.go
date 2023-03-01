package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"planigo/common"
	"planigo/internal/entities"
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

	assert.Equal(t, slots[0].Start, exceptedSlot[0].Start)
	assert.Equal(t, slots[0].End, exceptedSlot[0].End)
}

func TestGetNextDaysDate(t *testing.T) {

	dates := utils.GetNextDaysDate(2)

	expectedDates := []string{
		dateNow.Format("2006-01-02"),
		dateNow.AddDate(0, 0, 1).Format("2006-01-02"),
		dateNow.AddDate(0, 0, 2).Format("2006-01-02"),
	}

	assert.Equal(t, dates[0].Format("2006-01-02"), expectedDates[0])
	assert.Equal(t, dates[1].Format("2006-01-02"), expectedDates[1])
	assert.Equal(t, dates[2].Format("2006-01-02"), expectedDates[2])

	dates = utils.GetNextDaysDate(3)

	expectedDates = append(expectedDates, dateNow.AddDate(0, 0, 3).Format("2006-01-02"))

	assert.Equal(t, dates[3].Format("2006-01-02"), expectedDates[3])
}

func TestCreateEmptySlotsMapByShopHours(t *testing.T) {
	shopHoursByWeekDay := []entities.Hour{
		{Id: "1", Start: "09:00:00", End: "18:00:00", Day: 1, ShopID: "1"},
		{Id: "2", Start: "09:00:00", End: "18:00:00", Day: 2, ShopID: "1"},
		{Id: "3", Start: "09:00:00", End: "18:00:00", Day: 3, ShopID: "1"},
		{Id: "4", Start: "09:00:00", End: "18:00:00", Day: 4, ShopID: "1"},
		{Id: "5", Start: "09:00:00", End: "18:00:00", Day: 5, ShopID: "1"},
		{Id: "6", Start: "09:00:00", End: "18:00:00", Day: 6, ShopID: "1"},
		{Id: "7", Start: "09:00:00", End: "18:00:00", Day: 7, ShopID: "1"},
	}

	emptySlotsMap := utils.CreateEmptySlotsWithShopHours(shopHoursByWeekDay, 6)

	assert.Equal(t, len(emptySlotsMap), 7)
	assert.Equal(t, emptySlotsMap[0].Date, time.Now().Format("2006-01-02"))
	assert.Equal(t, len(emptySlotsMap[0].Slots), 9)
	assert.Equal(t, emptySlotsMap[1].Date, tomorrow.Format("2006-01-02"))
}

func TestMakeReservationMap(t *testing.T) {

	reservations := []common.DetailedReservation{
		{
			ReservationId: "1",
			ServiceId:     "1",
			ServiceName:   "Coiffure Homme",
			Start:         tomorrow.Format("2006-01-02 15:04:05"),
		},
	}

	reservationMap := utils.MakeReservationMap(reservations)

	key := fmt.Sprintf("%s-%s", tomorrow.Format("2006-01-02"), tomorrow.Format("15:04:05"))

	expectedReservationsMap := map[string]common.DetailedReservation{
		key: reservations[0],
	}

	assert.Equal(t, reservationMap[key].ReservationId, expectedReservationsMap[key].ReservationId)
}

func TestFillEmptySlotsWithReservationByDate(t *testing.T) {
	shopHoursByWeekDay := []entities.Hour{
		{Id: "1", Start: "11:00:00", End: "18:00:00", Day: 1, ShopID: "1"},
		{Id: "2", Start: "11:00:00", End: "18:00:00", Day: 2, ShopID: "1"},
		{Id: "3", Start: "11:00:00", End: "18:00:00", Day: 3, ShopID: "1"},
		{Id: "4", Start: "11:00:00", End: "18:00:00", Day: 4, ShopID: "1"},
		{Id: "5", Start: "11:00:00", End: "18:00:00", Day: 5, ShopID: "1"},
		{Id: "6", Start: "11:00:00", End: "18:00:00", Day: 6, ShopID: "1"},
		{Id: "7", Start: "11:00:00", End: "18:00:00", Day: 7, ShopID: "1"},
	}

	emptySlots := utils.CreateEmptySlotsWithShopHours(shopHoursByWeekDay, 6)

	daySlots := utils.FillEmptySlotsWithReservationByDate(emptySlots, []common.DetailedReservation{
		{ReservationId: "1", ServiceId: "1", ServiceName: "coiffure homme", Start: tomorrow.Format("2006-01-02 15:04:05")},
	})

	assert.False(t, daySlots[1].Slots[0].IsAvailable)
	assert.Equal(t, daySlots[1].Slots[0].ReservationId, "1")

}
