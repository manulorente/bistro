package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/lib/pq" //PostgreSQL Driver
)

var ormObject orm.Ormer

// ConnectToDb - Initializes the ORM and Connection to the postgres DB
func ConnectToDb() {
    envUser := os.Getenv("POSTGRES_USER")

    orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", "postgres", "user=dbUser password=yourPassword dbname=dbName host=dbHost sslmode=disable")
    orm.RegisterModel(new(User), new(Product))
    orm.DefaultTimeLoc = time.UTC
    ormObject = orm.NewOrm()
}

// GetOrmObject - Getter function for the ORM object with which we can query the database
func GetOrmObject() orm.Ormer {
    return ormObject
}