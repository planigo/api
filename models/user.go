package models

type User struct {
	Id              string `json:"id"`
	Email           string `json:"email"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Role            string `json:"role"`
	Password        string `json:"password,omitempty"`
	IsEmailVerified bool   `json:"isEmailVerified"`
}
