package userrepositorypostgresql

import (
	"fmt"

	userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"
	"github.com/darmayasa221/go-restapi/Infrastructures/database"
	"github.com/darmayasa221/go-restapi/Infrastructures/models"
)

type UserRepositoryPostgresql struct{}

func (u UserRepositoryPostgresql) GetUsers() interface{} {
	users := []map[string]interface{}{}
	database.DB.Table("users").Find(&users)
	return users
}

func (u UserRepositoryPostgresql) PostUser(p *userentities.User) *userentities.User {
	userModel := models.Users{User: *p}
	database.DB.Create(&userModel)
	return p
}

func (u UserRepositoryPostgresql) GetUser() {
	// var usr models.Users
	// var name string
	var us userentities.User
	database.DB.Model(&models.Users{}).Where("id = ?", 22).Select("name", "age").Scan(&us)
	fmt.Printf("%T", us)
}
