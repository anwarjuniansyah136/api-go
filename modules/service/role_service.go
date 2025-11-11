package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleService interface {
	Create(ctx *gin.Context)
}

type roleService struct {
	repository repository.RoleRepository
}

func NewRoleService(db *gorm.DB) RoleService {
	return &roleService{
		repository: repository.NewroleRepository(db),
	}
}

func (r *roleService) Create(ctx *gin.Context) {
	panic("unimplemented")
}