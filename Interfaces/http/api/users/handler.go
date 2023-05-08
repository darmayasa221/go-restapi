package users

import (
	"fmt"

	userusecase "github.com/darmayasa221/go-restapi/Applications/UseCase/user"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	uc userusecase.UserUseCase
}

func (uh UserHandlers) GetUserHandler(c *gin.Context) {
	fmt.Println("get user Handler")
	users := uh.uc.GetUserUseCase()
	c.JSON(200, gin.H{
		"message": "success user get user handler",
		"data":    users,
	})
}

func (uh UserHandlers) PostUserHandler(c *gin.Context) {
	fmt.Println("post user Handler")
	user := uh.uc.PostUserUseCase()
	c.JSON(201, gin.H{
		"message": "success user post user handler",
		"data":    user,
	})
}
