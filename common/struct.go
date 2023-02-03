package common

type DaySlot struct {
	Date  string `json:"date"`
	Slots []Slot `json:"slots"`
}

type Slot struct {
	ReservationId string `json:"reservationId"`
	IsAvailable   bool   `json:"isAvailable"`
	Start         string `json:"start"`
	End           string `json:"end"`
	Duration      int    `json:"duration"`
}

type DetailledReservation struct {
	ReservationId string `json:"reservationId"`
	ServiceId     string `json:"serviceId"`
	ServiceName   string `json:"serviceName"`
	Duration      string `json:"duration"`
	Start         string `json:"start"`
	End           string `json:"end"`
}
