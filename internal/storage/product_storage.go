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

func (s *ProductStorage) GetProducts() ([]models.Product, error) {
	query := `
	SELECT id, title, price, in_stock, created_at
	FROM products
	ORDER BY id ASC
	`
	rows, err := s.DB.Query(query)
	if err != nil {
		return []models.Product{}, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)

	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.ID,
			&product.Title,
			&product.Price,
			&product.InStock,
			&product.CreatedAt,
		)
		if err != nil {
			return []models.Product{}, err
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

func (s *ProductStorage) GetProduct(id int) (models.Product, error) {
	var product models.Product
	query := `SELECT id, title, price, in_stock, created_at
	FROM products
	WHERE id = $1`

	err := s.DB.QueryRow(query, id).Scan(
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

func (s *ProductStorage) UpdateProduct(id int, product models.Product) (models.Product, error) {
	query := `
	UPDATE products
	SET title = $1, price = $2, in_stock = $3
	WHERE id = $4
	RETURNING id, title, price, in_stock, created_at
	`
	err := s.DB.QueryRow(query, product.Title, product.Price, product.InStock, id).Scan(
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
