package userrepositories

import userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"

type UserRepositories interface {
	GetUsers() interface{}
	PostUser(p userentities.User) userentities.User
}
