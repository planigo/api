package store

import "database/sql"

type ServiceStore struct {
	*sql.DB
}

func NewServiceStore(db *sql.DB) *ServiceStore {
	return &ServiceStore{
		db,
	}
}
