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

func (u UserRepositoriesPostgresql) PostUser(p userentities.User) userentities.User {
	userModel := models.Users{User: p}
	database.DB.Create(&userModel)
	return p
}
