package userusecase

import (
	userrepositories "github.com/darmayasa221/go-restapi/Domain/user"
	userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"
)

type UserUseCase struct {
	userrepositories.UserRepositories
}

func (uc UserUseCase) GetUserUseCase() interface{} {
	users := uc.GetUser()
	return users
}

func (uc UserUseCase) PostUserUseCase() userentities.User {
	user := uc.PostUser()
	return user
}
