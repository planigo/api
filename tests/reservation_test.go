package tests

import (
	"planigo/utils"
	"testing"
)

func TestComputeEmptySlots(t *testing.T) {

	slots := utils.ComputeEmptySlots("09:00:00", "18:00:00")

	exceptedSlot := []utils.Slot{
		{
			Id:            "1",
			ReservationId: "",
			IsAvailable:   true,
			Start:         "09:00",
			End:           "10:00",
			Duration:      60,
		},
	}

	if slots[0] != exceptedSlot[0] {
		t.Errorf("got %s, want %s", slots[0].Start, exceptedSlot[0].Start)
	}
}
