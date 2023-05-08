package main

import (
	"github.com/darmayasa221/go-restapi/Infrastructures/database"
	"github.com/darmayasa221/go-restapi/Infrastructures/models"
)

func init() {
	database.ConnectToDB()
}

func main() {
	// migrate post data
	database.DB.AutoMigrate(&models.Users{})
}
