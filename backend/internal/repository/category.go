package repository

import (
	"database/sql"
	"good-and-new/internal/domain"
)

type CategoryRepository struct {
	db *sql.DB
}

type CategoryRepositoryInterface interface {
	FindAll() ([]domain.Category, error)
	Create(category domain.Category) (int64, error)
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) FindAll() ([]domain.Category, error) {
	var categories []domain.Category
	rows, err := cr.db.Query("select * from categories")
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var category domain.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return categories, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (cr *CategoryRepository) Create(category domain.Category) (int64, error) {
	stmt, err := cr.db.Prepare("insert into categories (name, created_at, updated_at) values(?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
