package handlers

import (
	// Third party
	"github.com/gin-gonic/gin"

)

func InitializeRoutes(r *gin.Engine) {

	// Setup end points
	r.GET("/", HomePage)
	//r.GET("/products", ProductsPage)
	//r.POST("/products", CreateProduct)
	r.POST("/query", QueryStrings)

	ur := r.Group("/user")
    {
        ur.GET("/register", ShowRegistrationPage)
		ur.POST("/register", Register)
        ur.GET("/login", ShowLoginPage)
        ur.POST("/login", PerformLogin)
		ur.GET("/logout", Logout)	

	}
	
	pr := r.Group("/products")
	{
		pr.GET("/view", ProductsPage)	
		pr.POST("/view", ProductsPage)	
		pr.POST("/create", CreateProduct)	
	}



	// Set new routes - One protected by Auth0 and other no
	//r.GET("/", handlers.HomePage)
	//r.GET("/products", handlers.AuthMiddleware(), handlers.ReadAllProducts)
	//r.POST("/products", handlers.AuthMiddleware(), handlers.CreateProduct)

}




