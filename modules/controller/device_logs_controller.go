package controller

import (
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeviceLogsController interface {
	Init()
}

type deviceLogController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewDeviceLogController(apiR *gin.RouterGroup, db *gorm.DB, ver string) DeviceLogsController {
	return &deviceLogController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (d *deviceLogController) Init() {
	deviceLogService := service.NewDeviceLogService(d.database)
	deviceLogs := d.apiRoutes.Group("/v1/device-logs")
	{
		deviceLogs.POST("/create", deviceLogService.Create)
		deviceLogs.GET("/find-all", deviceLogService.GetAllDevice)
	}
}