package controllers

import (
	// Third party
)


func (s *Server) InitializeRoutes() {

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	s.R.Use(SetUserStatus())
	s.R.LoadHTMLGlob("views/*")

	// Home route
	s.R.GET("/", HomePage)

	// USers routes
	ur := s.R.Group("/user/")
    {
		ur.GET("/register", EnsureNotLoggedIn(), ShowRegistrationPage)
		ur.GET("/login", EnsureNotLoggedIn(), ShowLoginPage)
		ur.GET("/logout", EnsureLoggedIn(), Logout)	
		ur.POST("/register", EnsureNotLoggedIn(), s.Register)
        ur.POST("/login", EnsureNotLoggedIn(), s.Login)
		//ur.GET("/:id/view", EnsureLoggedIn(), ProductsPage)	
		//ur.GET("/:id/edit", EnsureLoggedIn(), CreateProduct)
	}
	
	// Products routes
	pr := s.R.Group("/menu")
	{
		pr.POST("/", ViewMenu)
		//pr.GET("/view", EnsureLoggedIn(), ProductsPage)	
		//pr.GET("/view", EnsureLoggedIn(), ProductsPage)	
		//pr.POST("/view", EnsureLoggedIn(), ProductsPage)	
		//pr.POST("/create", EnsureLoggedIn(), CreateProduct)
		//pr.POST("/edit", EditProduct)	
	}

}




