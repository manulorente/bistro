package main

// En caso de error con alguna libreria instalada ejecutar go mod vendor
import (
	// System
    "io"
    "os"

	// Local
	"github.com/manulorente/bistro/handlers"

	// Third party
	"github.com/gin-gonic/gin"
)

// Port: Puerto de entrada al servidor
const (
	Port = ":3000"
)

var R *gin.Engine

// Punto de entrada del programa
func main() {

	// Set production mode or not
	gin.SetMode(gin.DebugMode)

	// Set the router as the default one shipped with Gin
	r := gin.Default()

	// Create log files
	logProd, err := os.Create("./logs/production.log")
	if err != nil {
		panic(err)
	}
	logErr, err := os.Create("./logs/error.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logProd)
	gin.DefaultErrorWriter = io.MultiWriter(logErr)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	r.LoadHTMLGlob("views/*")

	// Init routering
	handlers.InitializeRoutes(r)
	
	// Config and run the server
	r.Run(Port)
}
