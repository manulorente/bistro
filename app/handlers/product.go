package handlers

import (
	// System
	//"log"
	"net/http"

	// Local
	"github.com/manulorente/bistro/models"

	// Third party
	"github.com/gin-gonic/gin"
)

func ProductsPage(c *gin.Context) {
	products := models.GetAllProducts()
	Render(c, gin.H{
		"title": "Bistro",
		"msg":   "Menu de ejemplo",
		"payload":   products,
		},
		"menu.html")
}

// Function that fetches all products
func ReadAllProducts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, models.GetAllProducts())
}

/*
func CreateProduct(c *gin.Context) {
	var err error
	// Decodificamos el dato del cuerpo de la peticion y lo mapeamos a la estructura
	var product models.Product
	err = c.BindJSON(&product)
	if err != nil {
		log.Print(err)
	}

	models.ProductsList = append(models.ProductsList, product)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, product)

}*/

func CreateProduct(c *gin.Context) {
    cat := c.PostForm("cat")
    name := c.PostForm("name")
    description := c.PostForm("description")
	size := c.PostForm("size")    
	price := c.PostForm("price")
	
    if _, err := models.CreateNewProduct(cat, name, description, size, price); err == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/products/view")
    } else {
        c.AbortWithStatus(http.StatusBadRequest)
    }
}
