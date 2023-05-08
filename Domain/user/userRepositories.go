package userrepositories

import userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"

type UserRepositories interface {
	GetUser() interface{}
	PostUser() userentities.User
}
