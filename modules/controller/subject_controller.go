package controller

import (
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SubjectController interface {
	Init()
}

type subjectController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewSubjectController(apiR *gin.RouterGroup, db *gorm.DB, ver string) SubjectController {
	return &subjectController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *subjectController) Init() {
	subjectService := service.NewSubjectService(s.database)
	subject := s.apiRoutes.Group("/subject")
	{
		subject.POST("/create", subjectService.Create)
	}
}