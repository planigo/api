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

type DetailedReservation struct {
	ReservationId string `json:"reservationId"`
	ServiceId     string `json:"serviceId"`
	ServiceName   string `json:"serviceName"`
	Duration      string `json:"duration"`
	Start         string `json:"start"`
	End           string `json:"end"`
	UserId        string `json:"userId"`
}

type BookedReservation struct {
	ReservationId string  `json:"reservationId"`
	ShopName      string  `json:"shopName"`
	ServiceName   string  `json:"serviceName"`
	Duration      int32   `json:"duration"`
	Price         float32 `json:"price"`
	Start         string  `json:"start"`
	IsCancelled   bool    `json:"isCancelled"`
}

type AdminDetailedReservation struct {
	ReservationId string  `json:"reservationId"`
	Start         string  `json:"start"`
	Firstname     string  `json:"firstname"`
	Lastname      string  `json:"lastname"`
	ServiceName   string  `json:"serviceName"`
	Duration      int32   `json:"duration"`
	Price         float32 `json:"price"`
	IsCancelled   bool    `json:"isCancelled"`
}
