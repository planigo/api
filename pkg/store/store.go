package store

import (
	"database/sql"
	"planigo/internal/stores"
)

type Store struct {
	*stores.UserStore
	*stores.ServiceStore
	*stores.ShopStore
	*stores.HourStore
	*stores.CategoryStore
	*stores.ReservationStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		stores.NewUserStore(db),
		stores.NewServiceStore(db),
		stores.NewShopStore(db),
		stores.NewHourStore(db),
		stores.NewCategoryStore(db),
		stores.NewReservationStore(db),
	}
}
