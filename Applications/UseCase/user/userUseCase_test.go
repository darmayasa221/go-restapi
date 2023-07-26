package userusecase_test

import (
	"errors"
	"reflect"
	"testing"

	userusecase "github.com/darmayasa221/go-restapi/Applications/UseCase/user"
	userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"
)

type mockUserRepository struct{}

var count int = 0

var expectedUsers interface{} = []map[any]interface{}{
	{nil: userentities.User{Name: "expected name", Age: 1}},
}

func (m mockUserRepository) GetUsers() interface{} {
	count += 1
	return expectedUsers
}

func (m mockUserRepository) PostUser(p *userentities.User) *userentities.User {
	return p
}

func (m mockUserRepository) GetUser() {

}

func TestUserUseCase(t *testing.T) {
	t.Run("get users usecase", func(t *testing.T) {
		// Arrange
		/* initialization usecase instace */
		userUseCase := userusecase.UserUseCase{
			UserRepository: mockUserRepository{},
		}
		// Action
		result := userUseCase.GetUsersUseCase()
		/* must to be call mock GetUsers functin just one time */
		if r := 1; r != count {
			t.Errorf("getUser invoke over 1 time")
		}
		/* result at from usecase*/
		if !reflect.DeepEqual(result, expectedUsers) {
			t.Errorf("result not match")
		}
	})
	t.Run("post user usecase", func(t *testing.T) {
		t.Run("get error when payload not contain needed property", func(t *testing.T) {
			// Arrange
			/* mock validationPayload this can return error
			usecase use approach dependency inversion when use entities*/
			userentities.U.Age = 10
			expectedError := errors.New("not contain needed property")
			/* initialization usecase instace */
			userUseCase := userusecase.UserUseCase{
				UserRepository: mockUserRepository{},
			}
			// Action
			_, err := userUseCase.PostUserUseCase(userentities.U)
			// Assert
			if errors.Is(err, expectedError) {
				t.Error("should return error")
			}
		})
		t.Run("should return object corectly", func(t *testing.T) {
			// Arrange
			userentities.U = &userentities.User{
				Name: "expected name",
				Age:  1,
			}
			/* mock validationPayload this can return error
			usecase use approach dependency inversion when use entities*/
			expectedError := errors.New("not contain needed property")
			/* initialization usecase instace */
			userUseCase := userusecase.UserUseCase{
				UserRepository: mockUserRepository{},
			}
			// Action
			_, err := userUseCase.PostUserUseCase(userentities.U)
			// Assert
			if errors.Is(err, expectedError) {
				t.Error("should return error")
			}
		})
	})
}
