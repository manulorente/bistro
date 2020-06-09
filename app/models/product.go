package models

import (
	"time"
)

// Product : La API REST devolverá contenido JSON que permitirá que el frontend sea independiente al backend
type Product struct {
	ID          int       `json:"id" binding:"required"`
	Cat         string    `json:"cat,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Size        string    `json:"size,omitempty"`
	Price       string    `json:"price,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// BBDD virtual
var ProductsList = []Product{
	{ID: 0, Cat: "Comida", Name: "Caracoles", Description: "De Triana", Size: "Tarrina", Price: "5€", CreatedAt: time.Now()},
	{ID: 1, Cat: "Bebida", Name: "Vino", Description: "Rioja", Size: "Botella", Price: "12€", CreatedAt: time.Now()},
	{ID: 2, Cat: "Postre", Name: "Ensalada", Description: "De fruta", Size: "Unidad", Price: "5€", CreatedAt: time.Now()},
}

// Return a list of all the articles
func GetAllProducts() []Product {
	return ProductsList
}

func CreateNewProduct(cat, name, description, size, price string) (*Product, error) {
    p := Product{
		ID: len(ProductsList) + 1, 
		Cat : cat,
		Name: name,
		Description: description,
		Size: size,
		Price: price + "€",
		CreatedAt: time.Now(),
	}

    ProductsList = append(ProductsList, p)

    return &p, nil
}