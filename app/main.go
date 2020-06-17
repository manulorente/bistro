package main

// En caso de error con alguna libreria instalada ejecutar go mod vendor
import (
	// System
    //"io"
    "os"

	// Local
	models "github.com/manulorente/bistro/models"
	controllers "github.com/manulorente/bistro/controllers"

	// Third party
	//"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

)

var server = controllers.Server{}

func main() {

    if err := godotenv.Load(); err != nil {
        panic("No .env file found")
    }

	server.Initialize("postgres", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DB"))
	defer server.DB.Close()

	models.Load(server.DB)
	
	server.Run(os.Getenv("SERVER_PORT"))
}
