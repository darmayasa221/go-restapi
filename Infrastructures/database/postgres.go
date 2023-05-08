package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := "host=localhost user=developer password=developer dbname=gorestapi port=5432 "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		log.Fatal("Failed to connect database")
	}
}
