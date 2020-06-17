package controllers

import (
	"fmt"
	"log"
	"os"
	"io"

	"github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	
	// Local
	models "github.com/manulorente/bistro/models"
)

type Server struct {
	DB     	*gorm.DB
	R 		*gin.Engine
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	Dbdriver = "postgres"

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database ", Dbdriver)
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("We are connected to the %s database ", Dbdriver)
	}
	
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Product{}) //database migration

	// Set the router as the default one shipped with Gin
	server.R = gin.Default()
	server.R.Use(gin.Logger())
	server.R.Use(gin.Recovery())

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

	// Init routering
	server.InitializeRoutes()
	
}

func (server *Server) Run(addr string) {
	server.R.Run(":"+addr)
}