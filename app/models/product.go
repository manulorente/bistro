package models

import "time"

// Product : La API REST devolverá contenido JSON que permitirá que el frontend sea independiente al backend
type Product struct {
	ID          int       `json:"id" binding:"required"`
	Cat         string    `json:"cat,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Size        string    `json:"size,omitempty"`
	Price       string    `json:"price,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
