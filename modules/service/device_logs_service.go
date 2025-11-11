package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeviceLogsService interface {
	Create(ctx *gin.Context)
}

type devivceLogsService struct {
	repository repository.DeviceLogsRepository
}

func NewDeviceLogService(db *gorm.DB) DeviceLogsService {
	return &devivceLogsService{
		repository: repository.NewDeviceLogsRepository(db),
	}
}

func (d *devivceLogsService) Create(ctx *gin.Context) {
	panic("unimplemented")
}