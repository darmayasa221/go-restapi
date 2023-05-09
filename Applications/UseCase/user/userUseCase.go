package userusecase

import (
	userrepository "github.com/darmayasa221/go-restapi/Domain/user"
	userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"
)

type UserUseCase struct {
	userrepository.UserRepository
}

func (uc UserUseCase) GetUserUseCase() interface{} {
	users := uc.GetUsers()
	return users
}

func (uc UserUseCase) PostUserUseCase(p *userentities.User) (*userentities.User, error) {
	// validation payload
	err := userentities.ValidationPayload()

	if err != nil {
		return p, err
	}
	// validation success
	user := uc.PostUser(p)
	return user, nil
}
