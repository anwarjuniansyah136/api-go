package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SchoolService interface {
	Create(ctx *gin.Context)
}

type schoolService struct {
	repository repository.SchoolRepository
}

func NewSchoolServie(db *gorm.DB) SchoolService {
	return &schoolService{
		repository: repository.NewSchoolRepository(db),
	}
}

func (s *schoolService) Create(ctx *gin.Context) {
	panic("unimplemented")
}