package container

import (
	userusecase "github.com/darmayasa221/go-restapi/Applications/UseCase/user"
	userrepositoriespostgresql "github.com/darmayasa221/go-restapi/Infrastructures/repositories/postgresql"
)

type container struct {
	UserUseCase userusecase.UserUseCase
}

var Container container

func CreateContainer() {
	// initialization service locator
	userRepositoriesPostgresql := &userrepositoriespostgresql.UserRepositoriesPostgresql{}
	// usecase
	userUseCase := &userusecase.UserUseCase{UserRepositories: userRepositoriesPostgresql}
	// initialization on Contianer
	Container.UserUseCase = *userUseCase
}
