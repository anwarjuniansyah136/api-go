package controller

import (
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AttendanceRecordController interface {
	Init()
}

type attendaceRecordController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewAttendanceRecordController(apiR *gin.RouterGroup, db *gorm.DB, ver string) AttendanceRecordController {
	return &attendaceRecordController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (a *attendaceRecordController) Init() {
	attendanceRecordService := service.NewAttendanceRecordService(a.database)
	attendanceRecord := a.apiRoutes.Group("/v1/attendance-record")
	{
		attendanceRecord.POST("/create", attendanceRecordService.Create)
		attendanceRecord.GET("/find-all", attendanceRecordService.GetAllAttendanceRecord)
	}
}