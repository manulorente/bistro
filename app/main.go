package main

// En caso de error con alguna libreria instalada ejecutar go mod vendor
import (
	// System

	"log"

	// Local
	"github.com/manulorente/bistro/handlers"

	// Third party

	//"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//"github.com/rs/cors"
)

// Port: Puerto de entrada al servidor
const (
	Port = ":3000"
)

var R *gin.Engine

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
	r := gin.Default()

	// Serve frontend static files
	//r.Use(static.Serve("/", static.LocalFile("./static", true)))
	//r.Use(static.Serve("/", static.LocalFile("./views", true)))
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	r.LoadHTMLGlob("views/*")

	// Init routering
	handlers.InitializeRoutes(r)

	// For dev only - We use CORS to allow:
	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://google.com"}
	//config.AllowHeaders = []string{"Content-Type", "Origin", "Accept", "*"}
	//config.AllowCredentials = true

	// Config and run the server
	//r.Use(cors.New(config))
	r.Run(Port)
}
