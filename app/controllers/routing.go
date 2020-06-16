package controllers

import (
	// Third party
	"github.com/gin-gonic/gin"
)


func InitializeRoutes(r *gin.Engine) {

	// Setup end points
	r.Use(SetUserStatus())
	r.GET("/", HomePage)

	ur := r.Group("/user/")
    {
        ur.GET("/register", EnsureNotLoggedIn(), ShowRegistrationPage)
		ur.POST("/register", EnsureNotLoggedIn(), Register)
        ur.GET("/login", EnsureNotLoggedIn(), ShowLoginPage)
        ur.POST("/login", EnsureNotLoggedIn(), PerformLogin)
		ur.GET("/logout", EnsureLoggedIn(), Logout)	
		ur.GET("/:id/view", EnsureLoggedIn(), ProductsPage)	
		ur.GET("/:id/edit", EnsureLoggedIn(), CreateProduct)
	}
	
	pr := r.Group("/menu")
	{
		pr.POST("/", ViewMenu)
		//pr.GET("/view", EnsureLoggedIn(), ProductsPage)	
		//pr.GET("/view", EnsureLoggedIn(), ProductsPage)	
		//pr.POST("/view", EnsureLoggedIn(), ProductsPage)	
		//pr.POST("/create", EnsureLoggedIn(), CreateProduct)
		//pr.POST("/edit", EditProduct)	
	}

}




