package users

import (
	"fmt"

	userusecase "github.com/darmayasa221/go-restapi/Applications/UseCase/user"
	userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	uc *userusecase.UserUseCase
}

func (h UserHandlers) GetUserHandler(c *gin.Context) {
	users := h.uc.GetUsersUseCase()
	c.JSON(200, gin.H{
		"message": "success user get user handler",
		"data":    users,
	})
}

func (h UserHandlers) PostUserHandler(c *gin.Context) {
	// get payload from client
	c.BindJSON(userentities.U)
	// invoke use case
	user, err := h.uc.PostUserUseCase(userentities.U)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "success user post user handler",
		"data":    *user,
	})
}
