package store

import "database/sql"

type Store struct {
	*UserStore
	*ServiceStore
	*ShopStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		NewUserStore(db),
		NewServiceStore(db),
		NewShopStore(db),
	}
}
