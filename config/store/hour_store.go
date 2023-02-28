package store

import (
	"database/sql"
	"planigo/pkg/entities"
)

type HourStore struct {
	*sql.DB
}

func newHourStore(db *sql.DB) *HourStore {
	return &HourStore{
		db,
	}
}

func (s HourStore) GetHours() ([]entities.Hour, error) {
	var hours []entities.Hour

	query := "SELECT id, start, end, day, shop_id FROM Hour"

	rows, err := s.Query(query)
	if err != nil {
		return hours, err
	}

	for rows.Next() {
		hour := entities.Hour{}

		if err := rows.Scan(&hour.Id, &hour.Start, &hour.End, &hour.Day, &hour.ShopID); err != nil {
			return hours, err
		}

		hours = append(hours, hour)
	}

	return hours, nil
}

func (s HourStore) FindHoursByShopId(shopId string) ([]entities.Hour, error) {
	var hours []entities.Hour

	query := "SELECT id, start, end, day, shop_id FROM Hour WHERE shop_id = ?;"

	rows, err := s.Query(query, shopId)
	if err != nil {
		return hours, err
	}

	for rows.Next() {
		hour := entities.Hour{}

		if err := rows.Scan(&hour.Id, &hour.Start, &hour.End, &hour.Day, &hour.ShopID); err != nil {
			return hours, err
		}

		hours = append(hours, hour)
	}

	return hours, nil
}

func (s HourStore) CreateHour(hour entities.Hour) (entities.Hour, error) {
	insertedHour := entities.Hour{}

	query := "INSERT INTO Hour (start, end, day, shop_id) VALUES (?, ?, ?, ?) RETURNING id, start, end, day, shop_id"

	if err := s.
		QueryRow(query, hour.Start, hour.End, hour.Day, hour.ShopID).
		Scan(&insertedHour.Id, &insertedHour.Start, &insertedHour.End, &insertedHour.Day, &insertedHour.ShopID); err != nil {
		return insertedHour, err
	}

	return insertedHour, nil
}

func (s HourStore) GetHourById(shopId string) (entities.Hour, error) {
	hour := entities.Hour{}

	query := "SELECT id, start, end, day, shop_id FROM Hour WHERE id = ?"

	if err := s.
		QueryRow(query, shopId).
		Scan(&hour.Id, &hour.Start, &hour.End, &hour.Day, &hour.ShopID); err != nil {
		return hour, err
	}

	return hour, nil
}

func (s HourStore) GetHourByShopId(shopId string) ([]entities.Hour, error) {
	var hours []entities.Hour

	query := "SELECT id, start, end, day, shop_id FROM Hour WHERE shop_id = ?"

	rows, err := s.Query(query, shopId)
	if err != nil {
		return hours, err
	}

	for rows.Next() {
		hour := entities.Hour{}
		if err := rows.Scan(&hour.Id, &hour.Start, &hour.End, &hour.Day, &hour.ShopID); err != nil {
			return hours, err
		}

		hours = append(hours, hour)
	}

	return hours, nil
}

func (s HourStore) GetHourByShopIdAndDay(shopId string, weekDay int) (entities.Hour, error) {
	var hour entities.Hour

	if err := s.
		QueryRow("SELECT id, start, end, day, shop_id FROM Hour WHERE shop_id = ? AND day = ?", shopId, weekDay).
		Scan(&hour.Id, &hour.Start, &hour.End, &hour.Day, &hour.ShopID); err != nil {
		return hour, err
	}

	return hour, nil
}

func (s HourStore) DeleteHour(id string) error {
	query := "DELETE FROM Hour WHERE id = ?"

	if _, err := s.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (s HourStore) UpdateHour(id string, hour entities.Hour) (entities.Hour, error) {
	updatedHour := entities.Hour{}

	queryUpdate := "UPDATE Hour SET start = ?, end = ?, day = ? WHERE id = ?"
	if _, err := s.Exec(queryUpdate, hour.Start, hour.End, hour.Day, id); err != nil {
		return updatedHour, err
	}

	query := "SELECT id, start, end, day, shop_id FROM Hour WHERE id = ?"

	if err := s.
		QueryRow(query, id).
		Scan(&updatedHour.Id, &updatedHour.Start, &updatedHour.End, &updatedHour.Day, &updatedHour.ShopID); err != nil {
		return updatedHour, err
	}

	return updatedHour, nil
}
