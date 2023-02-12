package entities

import "time"

type Reservation struct {
	Id         string    `json:"id" `
	Start      time.Time `json:"start" `
	ServiceID  string    `json:"service_id" `
	UserID     string    `json:"user_id" `
	IsCanceled bool      `json:"is_canceled" `
}
