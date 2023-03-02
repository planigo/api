package stores

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
	"planigo/internal/entities"
)

type ServiceStore struct {
	*sql.DB
}

func NewServiceStore(db *sql.DB) *ServiceStore {
	return &ServiceStore{
		db,
	}
}

func (store *ServiceStore) FindServices() ([]entities.Service, error) {
	var services []entities.Service

	rows, err := store.Query("SELECT * FROM Service")

	if err != nil {
		return services, err
	}

	for rows.Next() {
		var serviceRow entities.Service
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

func (store *ServiceStore) FindServiceById(serviceId string) (entities.Service, error) {
	var service entities.Service

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

func (store *ServiceStore) FindServicesByShopId(shopId string) ([]entities.Service, error) {
	var services []entities.Service

	rows, err := store.Query("SELECT * FROM Service WHERE shop_id = ?;", shopId)

	if err != nil {
		return services, err
	}

	for rows.Next() {
		var serviceRow entities.Service
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

	return services, nil
}

func (store *ServiceStore) AddService(newService entities.Service) (string, error) {
	insertedService := new(entities.Service)
	query := "INSERT INTO Service (name, description, price, duration, shop_id) VALUES (?, ?, ?, ?, ?) RETURNING id"
	if err := store.
		QueryRow(query, newService.Name, newService.Description, newService.Price, newService.Duration, newService.ShopID).
		Scan(&insertedService.Id); err != nil {
		return insertedService.Id, err
	}

	return insertedService.Id, nil
}

func (store *ServiceStore) UpdateService(serviceId string, editedService entities.Service) (string, error) {
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
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusNoContent, nil
}
