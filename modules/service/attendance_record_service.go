package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AttendanceRecordService interface {
	Create(ctx *gin.Context)
}

type attendanceRecordService struct {
	repository repository.AttendanceRecordRepository
}

func NewAttendanceRecordService(db *gorm.DB) AttendanceRecordService {
	return &attendanceRecordService{
		repository: repository.NewAttendanceRecordRepository(db),
	}
}

func (a *attendanceRecordService) Create(ctx *gin.Context) {
	panic("unimplemented")
}