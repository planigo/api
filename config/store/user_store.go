package store

import (
	"database/sql"
	"planigo/pkg/entities"
)

type UserStore struct {
	*sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

func (store *UserStore) FindUsers() ([]entities.User, error) {
	var users []entities.User

	rows, err := store.Query("SELECT id, firstname, lastname FROM User")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var userRow entities.User
		if err := rows.Scan(&userRow.Id, &userRow.Firstname, &userRow.Lastname); err != nil {
			return users, err
		}
		users = append(users, userRow)
	}

	if err := rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (store *UserStore) CreateUser(user entities.User) (int, error) {
	panic("implement me")
	//...
}

func (store *UserStore) DeleteUser(id int) error {
	panic("implement me")
	//...
}

func (store *UserStore) UpdateUserById(id int) error {
	panic("implement me")
	//...
}

func (store *UserStore) FindUserById(id int) error {
	panic("implement me")
	//...
}
