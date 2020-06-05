package main

// En caso de error con alguna libreria instalada ejecutar go mod vendor
import (
	// System
	"log"
	"net/http"

	// Local
	"github.com/manulorente/bistro/handlers"

	// Third party
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// Port: Puerto de entrada al servidor
const (
	Port = ":3000"
)

// Punto de entrada del programa
func main() {

	// Here we are loading in our .env file which will contain our Auth0 Client Secret and Domain
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// register our actual jwtMiddleware
	handlers.RegMiddleware()

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	// For dev only - We use CORS to allow:
	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	config.AllowHeaders = []string{"Content-Type", "Origin", "Accept", "*"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// Set new routes - One protected by Auth0 and other no
	api.GET("/products", handlers.AuthMiddleware(), handlers.ReadAllProducts)
	api.POST("/products", handlers.AuthMiddleware(), handlers.CreateProduct)
	api.GET("/products2", handlers.ReadAllProducts)

	// Config and run the server
	router.Run(Port)
}
