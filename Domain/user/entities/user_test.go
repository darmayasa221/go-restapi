package userentities_test

import (
	"errors"
	"fmt"
	"testing"

	userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"
)

func TestUserEntiti(t *testing.T) {
	t.Run("should return object corectly", func(t *testing.T) {
		// Arrange
		user := userentities.User{
			Name: "test",
			Age:  1,
		}
		userentities.U = &user
		// Action
		err := userentities.ValidationPayload()
		// Assert
		if err != nil {
			t.Error(fmt.Sprint(err))
		}
	})
	t.Run("should return error when name is empty", func(t *testing.T) {
		// Arrange
		user := userentities.User{
			Name: "",
			Age:  1,
		}
		userentities.U = &user
		// Action
		err := userentities.ValidationPayload()
		exError := errors.New("not contain needed property")
		fmt.Println(err, exError)
		// Asssert
		if errors.Is(err, exError) {
			t.Error("not catch error")
		}
	})

}
