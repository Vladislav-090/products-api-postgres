package models

import "time"

type Product struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	InStock   int     `json:"stock"`
	CreatedAt time.Time `json:"createdat"`
}