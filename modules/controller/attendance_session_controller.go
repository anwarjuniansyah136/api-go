package controller

import (
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AttendanceSessionController interface {
	Init()
}

type attendanceSessionController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewAttendanceSessionController(apiR *gin.RouterGroup, db *gorm.DB, ver string) AttendanceRecordController {
	return &attendanceSessionController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (a *attendanceSessionController) Init() {
	attendanceSessionService := service.NewAttendanceSessionService(a.database)
	attendanceSession := a.apiRoutes.Group("/v1/attendance-session")
	{
		attendanceSession.POST("/create", attendanceSessionService.Create)
	}
}