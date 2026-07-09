package models

import "time"

type Product struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	InStock   bool      `json:"instock"`
	CreatedAt time.Time `json:"createdat"`
}
