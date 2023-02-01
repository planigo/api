package store

import "database/sql"

type Store struct {
	*UserStore
	*ServiceStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		NewUserStore(db),
		NewServiceStore(db),
	}
}
