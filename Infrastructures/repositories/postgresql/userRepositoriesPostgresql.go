package userrepositoriespostgresql

import (
	userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"
	"github.com/darmayasa221/go-restapi/Infrastructures/database"
	"github.com/darmayasa221/go-restapi/Infrastructures/models"
)

type UserRepositoriesPostgresql struct{}

func (u UserRepositoriesPostgresql) GetUser() interface{} {
	users := []map[string]interface{}{}
	database.DB.Table("users").Find(&users)
	return users
}

func (u UserRepositoriesPostgresql) PostUser() userentities.User {
	us := userentities.User{
		Name: "cobak",
		Age:  20,
	}
	user := models.Users{User: us}
	database.DB.Create(&user)
	return us
}
