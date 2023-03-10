package stores

import (
	"database/sql"
	"errors"
	"fmt"
	"planigo/common"
	"planigo/internal/entities"
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

func (r ReservationStore) GetReservationsByShopId(id string) ([]common.DetailedReservation, error) {
	var reservationList []common.DetailedReservation

	query := "SELECT r.id, s.id , s.name, s.duration, r.start FROM Reservation r, Service s WHERE r.service_id = s.id AND s.shop_id = ? AND is_cancelled = FALSE;"
	rows, err := r.Query(query, id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// get next week date by day of week

	for rows.Next() {
		reservation := common.DetailedReservation{}
		startDate, _ := time.Parse("2006-01-02 15:04:05", reservation.Start)
		duration, _ := strconv.Atoi(reservation.Duration)
		reservation.End = startDate.Add(time.Duration(duration) * time.Minute).String()
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
	query := "INSERT INTO Reservation (service_id, user_id, start) VALUES (?, ?, ?)"
	r.QueryRow(query, serviceId, userId, start)

	q := "SELECT id FROM Reservation WHERE id = LAST_INSERT_ID()"
	if err := r.QueryRow(q).Scan(&reservation.Id); err != nil {
		return "", err
	}

	return reservation.Id, nil
}

func (r ReservationStore) GetReservationById(id string) (common.DetailedReservation, error) {
	var reservation common.DetailedReservation
	query := "SELECT r.id, s.id, r.user_id, s.name, s.duration, r.start FROM Reservation r, Service s WHERE r.id = ? AND is_cancelled = false;"
	err := r.
		QueryRow(query, id).
		Scan(
			&reservation.ReservationId,
			&reservation.ServiceId,
			&reservation.UserId,
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
) (common.DetailedReservation, error) {
	serviceReservation := common.DetailedReservation{}
	query := "SELECT r.id, s.id , s.name, s.duration, r.start FROM Reservation r, Service s WHERE s.shop_id = ? AND r.start = ? AND r.is_cancelled = false;"
	rows, err := r.Query(query, shopId, start)
	if rows.Next() {
		return serviceReservation, errors.New("The slot is no longer available")
	}
	if err != nil {
		fmt.Println("??a se chie dessus : :", err.Error())
		return serviceReservation, err
	}

	uuid, err := r.InsertReservation(serviceId, start, userId)
	if err != nil {
		return serviceReservation, err
	}

	insertedReservation, _ := r.GetReservationById(uuid)

	return insertedReservation, nil
}

func (r ReservationStore) CancelReservation(id string) error {
	_, err := r.Exec("UPDATE Reservation SET is_cancelled = true WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r ReservationStore) GetSlotsBookedByUserId(userId string) ([]common.BookedReservation, error) {
	var reservationBooked []common.BookedReservation
	query := "SELECT r.id, sh.name, s.name, s.price, s.duration, r.start, r.is_cancelled FROM Reservation r JOIN Service s on s.id = r.service_id JOIN Shop sh on sh.id = s.shop_id WHERE r.user_id = ? AND r.is_cancelled = false ORDER BY r.`start` ASC"
	rows, err := r.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var reservation common.BookedReservation
		err := rows.
			Scan(
				&reservation.ReservationId,
				&reservation.ShopName,
				&reservation.ServiceName,
				&reservation.Price,
				&reservation.Duration,
				&reservation.Start,
				&reservation.IsCancelled,
			)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		reservationBooked = append(reservationBooked, reservation)
	}

	return reservationBooked, nil
}

func (r ReservationStore) FindSlotsBookedFilteredShopId(shopId string) ([]common.AdminDetailedReservation, error) {
	var reservations []common.AdminDetailedReservation
	query := "SELECT r.id, u.firstname, u.lastname, r.`start`, s.name, s.price, s.duration, r.is_cancelled FROM Reservation r JOIN `User` u ON u.id = r.user_id  JOIN Service s on s.id = r.service_id WHERE s.shop_id = ? AND r.is_cancelled = false"
	rows, err := r.Query(query, shopId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var reservation common.AdminDetailedReservation
		err := rows.
			Scan(
				&reservation.ReservationId,
				&reservation.Firstname,
				&reservation.Lastname,
				&reservation.Start,
				&reservation.ServiceName,
				&reservation.Price,
				&reservation.Duration,
				&reservation.IsCancelled,
			)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
