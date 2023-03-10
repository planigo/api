package stores

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
	"planigo/internal/entities"
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

	rows, err := store.Query("SELECT id, slug, name, description, owner_id, category_id FROM Shop")
	if err != nil {
		return shops, err
	}

	for rows.Next() {
		var shopRow entities.Shop
		if err := rows.Scan(&shopRow.Id, &shopRow.Slug, &shopRow.Name, &shopRow.Description, &shopRow.OwnerID, &shopRow.CategoryID); err != nil {
			return shops, err
		}
		shops = append(shops, shopRow)
	}

	if err := rows.Err(); err != nil {
		return shops, err
	}
	return shops, nil
}

func (store *ShopStore) FindShopsByUserId(userId string) ([]entities.Shop, error) {
	var shops []entities.Shop

	rows, err := store.Query("SELECT id, slug, name, description, owner_id, category_id FROM Shop WHERE owner_id = ?", userId)
	if err != nil {
		return shops, err
	}

	for rows.Next() {
		var shopRow entities.Shop
		if err := rows.Scan(&shopRow.Id, &shopRow.Slug, &shopRow.Name, &shopRow.Description, &shopRow.OwnerID, &shopRow.CategoryID); err != nil {
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

	row := store.QueryRow("SELECT id, slug, name, description, owner_id, category_id FROM Shop WHERE id = ?;", shopId)

	if err := row.Scan(&shop.Id, &shop.Slug, &shop.Name, &shop.Description, &shop.OwnerID, &shop.CategoryID); err != nil {
		return shop, err
	}

	return shop, nil
}

func (store *ShopStore) AddShop(newShop entities.ShopRequest) (entities.Shop, error) {
	insertedShop := new(entities.Shop)

	query := "INSERT INTO Shop (name, description, owner_id, category_id) VALUES (?, ?, ?, ?) RETURNING id, slug, name, description, category_id"

	if err := store.QueryRow(query, newShop.Name, newShop.Description, newShop.OwnerID, newShop.CategoryID).Scan(&insertedShop.Id, &insertedShop.Slug, &insertedShop.Name, &insertedShop.Description, &insertedShop.CategoryID); err != nil {
		return *insertedShop, err
	}

	return *insertedShop, nil
}

func (store *ShopStore) UpdateShop(shopId string, shopEdited entities.ShopRequest) (string, error) {
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
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusNoContent, nil
}

func (store *ShopStore) FindShopsByCategorySlug(categorySlug string) ([]entities.Shop, error) {
	var shops []entities.Shop

	rows, err := store.Query("SELECT s.id, s.slug, s.name, s.description, s.owner_id, s.category_id FROM Shop s LEFT JOIN Category ON Category.id = s.category_id WHERE Category.slug = ?;", categorySlug)
	if err != nil {
		return shops, err
	}

	for rows.Next() {
		var shopRow entities.Shop
		if err := rows.Scan(&shopRow.Id, &shopRow.Slug, &shopRow.Name, &shopRow.Description, &shopRow.OwnerID, &shopRow.CategoryID); err != nil {
			return shops, err
		}
		shops = append(shops, shopRow)
	}

	if err := rows.Err(); err != nil {
		return shops, err
	}
	return shops, nil
}
