package store

import (
	"database/sql"
	"log"
	"net/http"
	"planigo/pkg/entities"
)

type ShopStore struct {
	*sql.DB
}

func NewShopStore(db *sql.DB) *ShopStore {
	return &ShopStore{
		db,
	}
}

func (store *ShopStore) FindShops() ([]entities.Shop, error) {
	var shops []entities.Shop

	rows, err := store.Query("SELECT * FROM Shop")
	if err != nil {
		return shops, err
	}

	for rows.Next() {
		var shopRow entities.Shop
		if err := rows.Scan(&shopRow.Id, &shopRow.Name, &shopRow.Description, &shopRow.OwnerID); err != nil {
			return shops, err
		}
		shops = append(shops, shopRow)
	}

	if err := rows.Err(); err != nil {
		return shops, err
	}
	return shops, nil
}

func (store *ShopStore) FindShopById(shopId string) (entities.Shop, error) {
	var shop entities.Shop

	row := store.QueryRow("SELECT * FROM Shop WHERE id = ?;", shopId)

	if err := row.Scan(&shop.Id, &shop.Name, &shop.Description, &shop.OwnerID); err != nil {
		return shop, err
	}

	return shop, nil
}

func (store *ShopStore) UpdateShop(shopId string, shopEdited entities.ShopRequest) (entities.Shop, error) {
	row, err := store.Exec("UPDATE Shop SET name = ?, description = ? WHERE id = ?;", shopEdited.Name, shopEdited.Description, shopId)

	if err != nil {
		var shop entities.Shop
		return shop, err
	}

	if _, err = row.RowsAffected(); err != nil {
		log.Fatal(err)
	}

	shop, _ := store.FindShopById(shopId)

	return shop, nil
}

func (store *ShopStore) RemoveShop(shopId string) (int, error) {
	_, err := store.Exec("DELETE FROM Shop WHERE id = ?;", shopId)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
