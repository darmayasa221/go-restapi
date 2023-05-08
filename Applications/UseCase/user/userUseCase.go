package userusecase

import (
	userrepositories "github.com/darmayasa221/go-restapi/Domain/user"
	userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"
)

type UserUseCase struct {
	userrepositories.UserRepositories
}

func (uc UserUseCase) GetUserUseCase() interface{} {
	users := uc.GetUsers()
	return users
}

func (uc UserUseCase) PostUserUseCase(p userentities.User) userentities.User {
	user := uc.PostUser(p)
	return user
}
