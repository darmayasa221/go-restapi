package main

import (
	userusecase "github.com/darmayasa221/go-restapi/Applications/UseCase/user"
	"github.com/darmayasa221/go-restapi/Infrastructures/database"
	"github.com/darmayasa221/go-restapi/Infrastructures/loadEnv"
	userrepositoriespostgresql "github.com/darmayasa221/go-restapi/Infrastructures/repositories/postgresql"
	"github.com/darmayasa221/go-restapi/Interfaces/http/api/users"
	"github.com/gin-gonic/gin"
)

// initialization
func init() {
	loadEnv.LoadEnv()
	database.ConnectToDB()
}

func main() {
	route := gin.Default()

	userApi := route.Group("/user")
	users.Routes(userApi, userusecase.UserUseCase{UserRepositories: userrepositoriespostgresql.UserRepositoriesPostgresql{}})

	route.Run()
}
