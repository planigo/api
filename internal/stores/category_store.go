package stores

import (
	"database/sql"
	"planigo/internal/entities"
)

type CategoryStore struct {
	*sql.DB
}

func NewCategoryStore(db *sql.DB) *CategoryStore {
	return &CategoryStore{
		db,
	}
}

func (s CategoryStore) GetCategories() ([]entities.Category, error) {
	var categories []entities.Category

	query := "SELECT id, name, slug FROM Category"

	rows, err := s.Query(query)
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		category := entities.Category{}

		if err := rows.Scan(&category.Id, &category.Name, &category.Slug); err != nil {
			return categories, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
