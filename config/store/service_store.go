package store

import (
	"database/sql"
	"log"
	"net/http"
	"planigo/models"
)

type ServiceStore struct {
	*sql.DB
}

func NewServiceStore(db *sql.DB) *ServiceStore {
	return &ServiceStore{
		db,
	}
}

func (store *ServiceStore) FindServices() ([]models.Service, error) {
	var services []models.Service

	rows, err := store.Query("SELECT * FROM Service")

	if err != nil {
		return services, err
	}

	for rows.Next() {
		var serviceRow models.Service
		if err := rows.Scan(
			&serviceRow.Id,
			&serviceRow.Name,
			&serviceRow.Description,
			&serviceRow.Price,
			&serviceRow.Duration,
			&serviceRow.ShopID); err != nil {
			return services, err
		}
		services = append(services, serviceRow)
	}

	if err := rows.Err(); err != nil {
		return services, err
	}

	return services, nil
}

func (store *ServiceStore) FindServiceById(serviceId string) (models.Service, error) {
	var service models.Service

	row := store.QueryRow("SELECT * FROM Service WHERE id = ?;", serviceId)

	if err := row.Scan(
		&service.Id,
		&service.Name,
		&service.Description,
		&service.Price,
		&service.Duration,
		&service.ShopID); err != nil {
		return service, err
	}

	return service, nil
}

func (store *ServiceStore) AddService(newService models.Service) (string, error) {
	insertedService := new(models.Service)
	query := "INSERT INTO Service (name, description, price, duration, shop_id) VALUES (?, ?, ?, ?, ?) RETURNING id"
	if err := store.
		QueryRow(query, newService.Name, newService.Description, newService.Price, newService.Duration, newService.ShopID).
		Scan(&insertedService.Id); err != nil {
		return insertedService.Id, err
	}

	return insertedService.Id, nil
}

func (store *ServiceStore) UpdateService(serviceId string, editedService models.Service) (string, error) {
	row, err := store.Exec("UPDATE Service SET name = ?, description = ?, price = ?, duration = ? WHERE id = ?;", editedService.Name, editedService.Description, editedService.Price, editedService.Duration, serviceId)

	if err != nil {
		return "", err
	}

	if _, err = row.RowsAffected(); err != nil {
		log.Fatal(err)
	}

	return serviceId, nil
}

func (store *ServiceStore) RemoveService(serviceId string) (int, error) {
	_, err := store.Exec("DELETE FROM Service WHERE id = ?;", serviceId)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
