package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	Create(ctx *gin.Context)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		repository: repository.NewUserRepository(db),
	}
}

func (u *userService) Create(ctx *gin.Context) {
	panic("unimplemented")
}