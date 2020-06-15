package main

// En caso de error con alguna libreria instalada ejecutar go mod vendor
import (
	// System
    "io"
    "os"

	// Local
	"github.com/manulorente/bistro/handlers"
	"github.com/manulorente/bistro/models"

	// Third party
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/astaxie/beego/orm"

)

var R *gin.Engine
var ORM orm.Ormer

// Punto de entrada del programa
func main() {

    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        panic("No .env file found")
    }

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

	// Set production mode or not
	gin.SetMode(gin.DebugMode)

	// Connects to database and ther the ORM object and stores it in the global variable ORM
	models.ConnectToDb()
    ORM = models.GetOrmObject()
    ORM.Debug = true

	// Set the router as the default one shipped with Gin
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	r.LoadHTMLGlob("views/*")

	// Init routering
	handlers.InitializeRoutes(r)
	
	// Config and run the server
	r.Run(":"+os.GetEnv("APP_PORT"))
}
