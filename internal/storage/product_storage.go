package storage

import (
	"database/sql"
	"product-api-postgres/internal/models"
)

type ProductStorage struct {
	DB *sql.DB
}

func NewProductStorage(db *sql.DB) *ProductStorage {
	return &ProductStorage{
		DB: db,
	}
}

func (s *ProductStorage) CreateProduct(product models.Product) (models.Product, error) {
	query := `
	INSERT INTO products(title, price)
	VALUES ($1, $2)
	RETURNING id, title, price, in_stock, created_at
	`

	err := s.DB.QueryRow(query, product.Title, product.Price).Scan(
		&product.ID,
		&product.Title,
		&product.Price,
		&product.InStock,
		&product.CreatedAt,
	)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
