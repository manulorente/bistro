package storage

import (
	// System

	"time"

	// Third party
	"github.com/manulorente/bistro/models"
)

// BBDD virtual
var Products = []models.Product{
	{ID: 0, Cat: "Comida", Name: "Caracoles", Description: "De Triana", Size: "Tarrina", Price: "5€", CreatedAt: time.Now()},
	{ID: 1, Cat: "Bebida", Name: "Vino", Description: "Rioja", Size: "Botella", Price: "12€", CreatedAt: time.Now()},
	{ID: 2, Cat: "Postre", Name: "Ensalada", Description: "De fruta", Size: "Unidad", Price: "5€", CreatedAt: time.Now()},
}
