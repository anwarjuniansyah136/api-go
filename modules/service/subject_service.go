package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SubjectService interface {
	Create(ctx *gin.Context)
}

type subjectService struct {
	repository repository.SubjectRepository
}

func NewSubjectService(db *gorm.DB) SubjectService {
	return &subjectService{
		repository: repository.NewSubjectRepository(db),
	}
}

func (s *subjectService) Create(ctx *gin.Context) {
	panic("unimplemented")
}