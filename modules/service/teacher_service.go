package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeacherService interface {
	Create(ctx *gin.Context)
	GetAllTeacher(ctx *gin.Context)
}

type teacherService struct {
	repository repository.TeacherRepository
}

func NewTeacherService(db *gorm.DB) TeacherService {
	return &teacherService{
		repository: repository.NewTeacherRepository(db),
	}
}

func (t *teacherService) Create(ctx *gin.Context) {
	panic("unimplemented")
}


func (t *teacherService) GetAllTeacher(ctx *gin.Context) {
	panic("unimplemented")
}