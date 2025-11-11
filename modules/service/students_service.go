package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentsService interface {
	Create(ctx *gin.Context)
}

type studentsService struct {
	repository repository.StudentRepository
}

func NewStudentsService(db *gorm.DB) StudentsService {
	return &studentsService{
		repository: repository.NewStudentRepository(db),
	}
}

func (s *studentsService) Create(ctx *gin.Context) {
	panic("unimplemented")
}