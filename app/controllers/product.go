package controllers

import (
	// System
	//"log"
	//"net/http"

	// Local
	//models "github.com/manulorente/bistro/models"

	// Third party
	//"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"

)
/*
func ProductsPage(c *gin.Context) {
	products := GetAllProducts()
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
	c.JSON(http.StatusOK, GetAllProducts())
}
*/
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
/*
func CreateProduct(c *gin.Context) {
    cat := c.PostForm("cat")
    name := c.PostForm("name")
    description := c.PostForm("description")
	size := c.PostForm("size")    
	price := c.PostForm("price")
	
    if _, err := CreateNewProduct(cat, name, description, size, price); err == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/products/view")
    } else {
        c.AbortWithStatus(http.StatusBadRequest)
    }
}

func GetAllProducts() (c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func CreateNewProduct(cat, name, description, size, price string) (*Product, error) {
    p := models.Product{
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
*/