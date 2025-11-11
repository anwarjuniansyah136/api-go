package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AttendanceSessionService interface {
	Create(ctx *gin.Context)
}

type attendanceSessionService struct {
	repository repository.AttendanceSessionRepository
}

func NewAttendanceSessionService(db *gorm.DB) AttendanceSessionService {
	return &attendanceSessionService{
		repository: repository.NewAttendanceSessionRepository(db),
	}
}

func (a *attendanceSessionService) Create(ctx *gin.Context) {
	panic("unimplemented")
}