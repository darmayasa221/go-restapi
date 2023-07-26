package userrepository

import userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"

type UserRepository interface {
	GetUsers() interface{}
	PostUser(p *userentities.User) *userentities.User
	GetUser()
}
