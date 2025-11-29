package controller

import (
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ScheduleClassController interface {
	Init()
}

type scheduleClassController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewScheduleClassController(apiR *gin.RouterGroup, db *gorm.DB, ver string) ScheduleClassController {
	return &scheduleClassController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *scheduleClassController) Init() {
	scheduleClassService := service.NewScheduleClassService(s.database)
	scheduleClass := s.apiRoutes.Group("/v1/schedule-class")
	{
		scheduleClass.POST("/create", scheduleClassService.Create)
		scheduleClass.GET("/find-all", scheduleClassService.GetAllScheduleClass)
	}
}