package store

import "database/sql"

type Store struct {
	*UserStore
	*ServiceStore
	*ShopStore
	*HourStore
	*CategoryStore
	*ReservationStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		NewUserStore(db),
		NewServiceStore(db),
		NewShopStore(db),
		newHourStore(db),
		newCategoryStore(db),
		NewReservationStore(db),
	}
}
