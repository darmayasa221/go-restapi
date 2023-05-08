package users

import (
	userusecase "github.com/darmayasa221/go-restapi/Applications/UseCase/user"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, uc userusecase.UserUseCase) {
	uh := UserHandlers{
		uc: uc,
	}
	r.GET("/", uh.GetUserHandler)
	r.POST("/", uh.PostUserHandler)
}
