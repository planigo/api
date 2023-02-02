package store

import (
	"database/sql"
	"fmt"
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

	rows, err := store.Query("SELECT id, firstname, lastname, role, is_email_verified FROM User")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var userRow entities.User
		if err := rows.Scan(&userRow.Id, &userRow.Firstname, &userRow.Lastname, &userRow.Role, &userRow.IsEmailVerified); err != nil {
			return users, err
		}
		users = append(users, userRow)
	}

	if err := rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (store *UserStore) CreateUser(user entities.User) (string, error) {
	query := "INSERT INTO User (id, email, firstname, lastname, role, password) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := store.DB.Exec(query, "", user.Email, user.Firstname, user.Lastname, user.Role, user.Password)
	if err != nil {
		return "", err
	}

	uuid, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	fmt.Println("User created with id: ", uuid)

	return string(uuid), nil
}

func (store *UserStore) FindUserByEmail(email string) (entities.User, error) {
	user := entities.User{}

	query := "SELECT id, email, firstname, lastname, role, password, is_email_verified FROM User WHERE email = ?"
	err := store.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &user.Role, &user.Password, &user.IsEmailVerified)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (store *UserStore) DeleteUser(id int) error {
	panic("implement me")
	//...
}

func (store *UserStore) UpdateUserById(id int) error {
	panic("implement me")
	//...
}

func (store *UserStore) FindUserById(id string) (entities.User, error) {
	user := &entities.User{}

	query := "SELECT id, email, firstname, lastname, role FROM User WHERE id = ?"
	err := store.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &user.Role)
	if err != nil {
		return *user, err
	}
	return *user, nil
}
