package store

import (
	"database/sql"
	"log"
	"net/http"
	"planigo/models"
)

type ShopStore struct {
	*sql.DB
}

func NewShopStore(db *sql.DB) *ShopStore {
	return &ShopStore{
		db,
	}
}

func (store *ShopStore) FindShops() ([]models.Shop, error) {
	var shops []models.Shop

	rows, err := store.Query("SELECT * FROM Shop")
	if err != nil {
		return shops, err
	}

	for rows.Next() {
		var shopRow models.Shop
		if err := rows.Scan(&shopRow.Id, &shopRow.Name, &shopRow.Description, &shopRow.OwnerID, &shopRow.CategoryID); err != nil {
			return shops, err
		}
		shops = append(shops, shopRow)
	}

	if err := rows.Err(); err != nil {
		return shops, err
	}
	return shops, nil
}

func (store *ShopStore) FindShopById(shopId string) (models.Shop, error) {
	var shop models.Shop

	row := store.QueryRow("SELECT * FROM Shop WHERE id = ?;", shopId)

	if err := row.Scan(&shop.Id, &shop.Name, &shop.Description, &shop.OwnerID, &shop.CategoryID); err != nil {
		return shop, err
	}

	return shop, nil
}

func (store *ShopStore) AddShop(newShop models.ShopRequest) (models.Shop, error) {
	insertedShop := new(models.Shop)

	query := "INSERT INTO Shop (name, description, owner_id, category_id) VALUES (?, ?, ?, ?) RETURNING id, name, description, category_id"

	if err := store.QueryRow(query, newShop.Name, newShop.Description, newShop.OwnerID, newShop.CategoryID).Scan(&insertedShop.Id, &insertedShop.Name, &insertedShop.Description, &insertedShop.CategoryID); err != nil {
		return *insertedShop, err
	}

	return *insertedShop, nil
}

func (store *ShopStore) UpdateShop(shopId string, shopEdited models.ShopRequest) (string, error) {
	row, err := store.Exec("UPDATE Shop SET name = ?, description = ?, category_id = ? WHERE id = ?;", shopEdited.Name, shopEdited.Description, shopEdited.CategoryID, shopId)

	if err != nil {
		return "", err
	}

	if _, err = row.RowsAffected(); err != nil {
		log.Fatal(err)
	}

	return shopId, nil
}

func (store *ShopStore) RemoveShop(shopId string) (int, error) {
	_, err := store.Exec("DELETE FROM Shop WHERE id = ?;", shopId)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
