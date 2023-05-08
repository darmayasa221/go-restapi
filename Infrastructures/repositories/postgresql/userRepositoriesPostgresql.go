package userrepositoriespostgresql

import (
	userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"
	"github.com/darmayasa221/go-restapi/Infrastructures/database"
	"github.com/darmayasa221/go-restapi/Infrastructures/models"
)

type UserRepositoriesPostgresql struct{}

func (u UserRepositoriesPostgresql) GetUsers() interface{} {
	users := []map[string]interface{}{}
	database.DB.Table("users").Find(&users)
	return users
}

func (u UserRepositoriesPostgresql) PostUser() userentities.User {
	userEntitis := userentities.User{
		Name: "cobak",
		Age:  20,
	}
	userModel := models.Users{User: userEntitis}
	database.DB.Create(&userModel)
	return userEntitis
}
