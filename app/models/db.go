package models

import (
    "time"
    "os"
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

// Initialise value
var sample = User{ 
        UserID: 0, Email: "admin@bistro.com", Username: "admin", Password: "admin", CreatedAt: time.Now(), UpdatedAt: time.Time{}, //DeletedAt: null,
     }

/*
var sample = User{ 
        UserID: 0, Email: "admin@bistro.com", Username: "admin", Password: "admin", CreatedAt: time.Now(), UpdatedAt: time.Time{},//DeletedAt: "-",
        Products: []Product{
            {ProductID: 0, Cat: "Comida", Name: "Caracoles", Description: "De Triana", Size: "Tarrina", Price: "5€", CreatedAt: time.Now()},
            {ProductID: 1, Cat: "Bebida", Name: "Vino", Description: "Rioja", Size: "Botella", Price: "12€", CreatedAt: time.Now()},
            {ProductID: 2, Cat: "Postre", Name: "Ensalada", Description: "De fruta", Size: "Unidad", Price: "5€", CreatedAt: time.Now()},            
        },
    }
*/

// ConnectToDb - Initializes the GORM and Connection to the postgres DB
func ConnectToDb() *gorm.DB {
    envHost := os.Getenv("POSTGRES_HOST")
    envPort := os.Getenv("POSTGRES_PORT")
    envUser := os.Getenv("POSTGRES_USER")
    envDb := os.Getenv("POSTGRES_DB")
    envPassword := os.Getenv("POSTGRES_PASSWORD")

    // https://gobyexample.com/string-formatting
    prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", 
    envHost, envPort, envUser, envDb, envPassword)
    fmt.Println("conname is\t\t", prosgret_conname)

    db, err := gorm.Open("postgres", prosgret_conname)
    if err != nil {
        panic("Failed to connect to database!")
    }

    db.AutoMigrate(&User{})

    db.Create(&sample)

    return db
}
