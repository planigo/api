package entities

import "time"

type Hours struct {
	Id      string    `json:"id"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Day     int       `json:"day"`
	ShopID string    `json:"store_id"`
}
