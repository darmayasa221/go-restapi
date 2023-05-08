package models

import (
	userentities "github.com/darmayasa221/go-restapi/Domain/user/entities"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	userentities.User
}
