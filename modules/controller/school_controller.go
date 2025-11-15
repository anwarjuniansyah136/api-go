package controller

import (
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SchoolController interface {
	Init()
}

type schoolController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewSchoolController(apiR *gin.RouterGroup, db *gorm.DB, ver string) SchoolController {
	return &schoolController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *schoolController) Init() {
	schoolService := service.NewSchoolServie(s.database)
	school := s.apiRoutes.Group("/v1/school")
	{
		school.POST("/create", schoolService.Create)
		school.GET("/find-all", schoolService.GetAllUser)
	}
}