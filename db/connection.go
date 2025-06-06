package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=admin password=ac040487 dbname=postgres port=5432"
var DB *gorm.DB
var err error

func DBConnect() {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database")
		panic("failed to connect to database")
	} else {
		log.Println("Connected to database successfully")
	}
}
