package container

import (
	userusecase "github.com/darmayasa221/go-restapi/Applications/UseCase/user"
	userrepositoryespostgresql "github.com/darmayasa221/go-restapi/Infrastructures/repository/postgresql"
)

type container struct {
	UserUseCase *userusecase.UserUseCase
}

var Container *container = &container{}

func CreateContainer() {
	// initialization service locator
	userRepositoryPostgresql := &userrepositoryespostgresql.UserRepositoryPostgresql{}
	// usecase
	userUseCase := &userusecase.UserUseCase{UserRepository: userRepositoryPostgresql}
	// initialization on Contianer
	Container.UserUseCase = userUseCase
}
