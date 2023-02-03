package store

import (
	"database/sql"
	"errors"
	"fmt"
	"planigo/common"
	"planigo/pkg/entities"
	"strconv"
	"time"
)

type ReservationStore struct {
	*sql.DB
}

func NewReservationStore(db *sql.DB) *ReservationStore {
	return &ReservationStore{
		db,
	}
}

func (r ReservationStore) GetReservationsByShopId(id string) ([]common.DetailledReservation, error) {
	var reservationList []common.DetailledReservation

	query := "SELECT r.id, s.id , s.name, s.duration, r.start FROM Reservation r, Service s WHERE r.service_id = s.id AND s.shop_id = ?;"
	rows, err := r.Query(query, id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// get next week date by day of week

	for rows.Next() {
		reservation := common.DetailledReservation{}
		startDate, _ := time.Parse("2006-01-02 15:04:05", reservation.Start)
		fmt.Println(startDate)
		duration, _ := strconv.Atoi(reservation.Duration)
		fmt.Println(duration)
		reservation.End = startDate.Add(time.Duration(duration) * time.Minute).String()
		fmt.Println(reservation.End)
		err := rows.
			Scan(
				&reservation.ReservationId,
				&reservation.ServiceId,
				&reservation.ServiceName,
				&reservation.Duration,
				&reservation.Start,
			)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		reservationList = append(reservationList, reservation)
	}
	return reservationList, nil

}

func (r ReservationStore) InsertReservation(serviceId string, start string, userId string) (string, error) {
	var reservation entities.Reservation
	query := "INSERT INTO Reservation (service_id, user_id, start) VALUES (?, ?, ?) RETURNING id;"
	err := r.QueryRow(query, serviceId, userId, start).Scan(&reservation.Id)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return reservation.Id, nil
}

func (r ReservationStore) GetReservationById(id string) (common.DetailledReservation, error) {
	var reservation common.DetailledReservation
	query := "SELECT r.id, s.id , s.name, s.duration, r.start FROM Reservation r, Service s WHERE r.id = ?;"
	err := r.
		QueryRow(query, id).
		Scan(
			&reservation.ReservationId,
			&reservation.ServiceId,
			&reservation.ServiceName,
			&reservation.Duration,
			&reservation.Start,
		)
	if err != nil {
		fmt.Println(err.Error())
		return reservation, err
	}

	return reservation, nil
}

func (r ReservationStore) BookReservation(
	serviceId string,
	shopId string,
	start string,
	userId string,
) (common.DetailledReservation, error) {
	serviceReservation := common.DetailledReservation{}
	query := "SELECT r.id, s.id , s.name, s.duration, r.start FROM Reservation r, Service s WHERE s.shop_id = ? AND r.start = ?;"
	rows, err := r.Query(query, shopId, start)
	if rows.Next() {
		return serviceReservation, errors.New("Reservation already booked")
	}
	if err != nil {
		fmt.Println("ça se chie dessus : :", err.Error())
		return serviceReservation, err
	}

	uuid, err := r.InsertReservation(serviceId, start, userId)
	if err != nil {
		return serviceReservation, err
	}

	insertedReservation, _ := r.GetReservationById(uuid)

	return insertedReservation, nil
}
