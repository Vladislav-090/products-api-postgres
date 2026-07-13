package storage

import (
	"database/sql"
	"product-api-postgres/internal/models"
)

type UserStorage struct {
	DB *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		DB: db,
	}
}

func (s *UserStorage) CreateUser(user models.User) (models.User, error) {
	query := `
	INSERT INTO users (email, password_hash, role)
	VALUES ($1, $2, $3)
	RETURNING id, email, password_hash, role, created_at
	`

	err := s.DB.QueryRow(
		query,
		user.Email,
		user.PasswordHash,
		user.Role,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *UserStorage) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT id, email, password_hash, role, created_at
	FROM users
	WHERE email = $1`

	err := s.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
