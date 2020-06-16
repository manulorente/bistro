package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ProductID   int       	`json:"productid"`
	Cat         string   	`json:"cat,omitempty"`
	Name        string    	`json:"name" binding:"required"`
	Description string    	`json:"description,omitempty"`
	Size        string    	`json:"size" binding:"required"`
	Price       string    	`json:"price" binding:"required"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	DeletedAt   *time.Time  `json:"deleted_at,omitempty"`
}


/*
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
}*/