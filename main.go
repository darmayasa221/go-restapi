package main

import (
	container "github.com/darmayasa221/go-restapi/Infrastructures"
	"github.com/darmayasa221/go-restapi/Infrastructures/database"
	"github.com/darmayasa221/go-restapi/Infrastructures/loadEnv"
	"github.com/darmayasa221/go-restapi/Interfaces/http/api/users"

	"github.com/gin-gonic/gin"
)

// initialization
func init() {
	loadEnv.LoadEnv()
	database.ConnectToDB()
	container.CreateContainer()
}

func main() {
	// service locator
	c := container.Container

	route := gin.Default()

	userApi := route.Group("/user")
	users.Routes(userApi, &c.UserUseCase)

	route.Run()
}
