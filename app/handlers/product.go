package handlers

import (
	// System
	"log"
	"net/http"

	// Local
	"github.com/manulorente/bistro/models"
	"github.com/manulorente/bistro/storage"

	// Third party
	"github.com/gin-gonic/gin"
)

func ReadAllProducts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, storage.Products)
}

func CreateProduct(c *gin.Context) {
	var err error
	// Decodificamos el dato del cuerpo de la peticion y lo mapeamos a la estructura
	var product models.Product
	err = c.BindJSON(&product)
	if err != nil {
		log.Print(err)
	}

	storage.Products = append(storage.Products, product)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, product)

}
