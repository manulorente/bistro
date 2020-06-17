package models

import (
    "time"
    "log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)


// Initialise value
var users = []User{ 
        User{
            Email: "admin@bistro.com", 
            Username: "admin", 
            Password: "admin", 
            CreatedAt: time.Now(), 
            UpdatedAt: time.Now(),
        },
        User{
            Email: "user1@gmail.com", 
            Username: "user1", 
            Password: "user1", 
            CreatedAt: time.Now(), 
            UpdatedAt: time.Now(),
        },
}


var products = []Product{ 
        Product{
            Category: "Comida", 
            Name: "Caracoles", 
            Description: "De Triana", 
            Size: "Tarrina", 
            Price: "5€", 
            CreatedAt: time.Now(), 
            UpdatedAt: time.Now(),
        },      
        Product{
            Category: "Bebida", 
            Name: "Vino", 
            Description: "Rioja", 
            Size: "Botella", 
            Price: "12€", 
            CreatedAt: time.Now(), 
            UpdatedAt: time.Now(),
        },
        Product{
            Category: "Postre", 
            Name: "Ensalada", 
            Description: "De fruta", 
            Size: "Unidad", 
            Price: "5€", 
            CreatedAt: time.Now(), 
            UpdatedAt: time.Now(),
        },            
}

/*
// ConnectToDb - Initializes the GORM and Connection to the postgres DB
func ConnectToDb() *gorm.DB {
    var err error

    envHost := os.Getenv("POSTGRES_HOST")
    envPort := os.Getenv("POSTGRES_PORT")
    envUser := os.Getenv("POSTGRES_USER")
    envDb := os.Getenv("POSTGRES_DB")
    envPassword := os.Getenv("POSTGRES_PASSWORD")

    // https://gobyexample.com/string-formatting
    prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", 
    envHost, envPort, envUser, envDb, envPassword)
    fmt.Println("conname is\t\t", prosgret_conname)

    DB, err := gorm.Open("postgres", prosgret_conname)
    if err != nil {
        panic("Failed to connect to database!")
    }
    fmt.Println("You are Successfully connected!")

    DB.Debug().AutoMigrate(&User{}, &Product{}) //database migration
    //db.Create(&sample)

    return DB
}*/

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&Product{}, &User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&User{}, &Product{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&Product{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		products[i].UserID = users[i].ID

		err = db.Debug().Model(&Product{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}