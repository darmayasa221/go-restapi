package users

import (
	"fmt"

	userusecase "github.com/darmayasa221/go-restapi/Applications/UseCase/user"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	uc *userusecase.UserUseCase
}

func (h UserHandlers) GetUserHandler(c *gin.Context) {
	fmt.Println("get user Handler")
	users := h.uc.GetUserUseCase()
	c.JSON(200, gin.H{
		"message": "success user get user handler",
		"data":    users,
	})
}

func (h UserHandlers) PostUserHandler(c *gin.Context) {
	fmt.Println("post user Handler")
	user := h.uc.PostUserUseCase()
	c.JSON(201, gin.H{
		"message": "success user post user handler",
		"data":    user,
	})
}
