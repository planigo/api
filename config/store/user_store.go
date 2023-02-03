package store

import (
	"database/sql"
	"planigo/models"
	"planigo/utils"
)

type UserStore struct {
	*sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

func (store *UserStore) FindUsers() ([]models.User, error) {
	var users []models.User

	rows, err := store.Query("SELECT id, firstname, lastname, role, is_email_verified FROM User")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var userRow models.User
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

func (store *UserStore) CreateUser(user models.User) (string, error) {
	insertedUser := &models.User{}

	query := "INSERT INTO User (email, firstname, lastname, role, password) VALUES (?, ?, ?, ?, ?) RETURNING id, email"

	if err := store.
		QueryRow(query, user.Email, user.Firstname, user.Lastname, user.Role, user.Password).
		Scan(&insertedUser.Id, &insertedUser.Email); err != nil {
		return "", err
	}

	return insertedUser.Id, nil
}

func (store *UserStore) FindUserByEmail(email string) (models.User, error) {
	user := models.User{}

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

func (store *UserStore) FindUserById(id string) (models.User, error) {
	user := &models.User{}

	query := "SELECT id, email, firstname, lastname, role, is_email_verified FROM User WHERE id = ?"
	err := store.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &user.Role, &user.IsEmailVerified)
	if err != nil {
		return *user, err
	}
	return *user, nil
}

func (store *UserStore) ValidateUserEmail(token string) error {

	payload, err := utils.VerifyJWT(token)
	if err != nil {
		println("Error: VerifyJWT ", err.Error(), "\n")
		return err
	}

	user := &models.User{}
	query := "SELECT id FROM User WHERE id = ? AND is_email_verified = 0"
	err = store.QueryRow(query, payload.ID).Scan(&user.Id)
	if err != nil {
		println("Error: QueryRow ", err, "\n")
		return err
	}

	println("User id: ", user.Id, "\n")

	_, err = store.Exec("UPDATE User SET `is_email_verified` = 1 WHERE id = ?", user.Id)
	if err != nil {
		println("Error: Exec ", err, "\n")
		return err
	}

	return nil
}
