package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ScheduleClassService interface {
	Create(ctx *gin.Context)
}

type scheduleClassService struct {
	repository repository.ScheduleClassRepository
}

func NewScheduleClassService(db *gorm.DB) ScheduleClassService {
	return &scheduleClassService{
		repository: repository.NewScheduleRepository(db),
	}
}

func (s *scheduleClassService) Create(ctx *gin.Context) {
	panic("unimplemented")
}