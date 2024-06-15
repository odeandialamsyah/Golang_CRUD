package entities

import "time"

type Product struct {
	Id          uint
	Name        string
	Category    category
	Stock       int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}