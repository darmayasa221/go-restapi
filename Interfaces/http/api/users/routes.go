package users

import (
	userusecase "github.com/darmayasa221/go-restapi/Applications/UseCase/user"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, uc *userusecase.UserUseCase) {
	// inject usescase at handler
	h := UserHandlers{uc: uc}
	//
	r.GET("/", h.GetUserHandler)
	r.POST("/", h.PostUserHandler)
}
