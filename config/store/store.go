package store

import "database/sql"

type Store struct {
	*UserStore
	*ServiceStore
	*ShopStore
	*HourStore
	*ReservationStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		NewUserStore(db),
		NewServiceStore(db),
		NewShopStore(db),
		newHourStore(db),
		NewReservationStore(db),
	}
}
