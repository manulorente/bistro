package handlers

import (
	// Third party
	"github.com/gin-gonic/gin"
	"regexp"

)

var Regex_var = regexp.MustCompile("^/(register|login|view|create)/([a-zA-Z0-9]+)$")

func InitializeRoutes(r *gin.Engine) {

	r.Use(SetUserStatus())

	// Setup end points
	r.GET("/", HomePage)
	r.POST("/query", QueryStrings)

	ur := r.Group("/user")
    {
        ur.GET("/register", EnsureNotLoggedIn(), ShowRegistrationPage)
		ur.POST("/register", EnsureNotLoggedIn(), Register)
        ur.GET("/login", EnsureNotLoggedIn(), ShowLoginPage)
        ur.POST("/login", EnsureNotLoggedIn(), PerformLogin)
		ur.GET("/logout", EnsureLoggedIn(), Logout)	
	}
	
	pr := r.Group("/products")
	{
		pr.GET("/view", EnsureLoggedIn(), ProductsPage)	
		pr.POST("/view", EnsureLoggedIn(), ProductsPage)	
		pr.POST("/create", EnsureLoggedIn(), CreateProduct)
		//pr.POST("/edit", EditProduct)	
	}

}




