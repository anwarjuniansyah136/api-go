package controller

import (
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentsController interface {
	Init()
}

type studentsController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}
func NewStudentsController(apiR *gin.RouterGroup, db *gorm.DB, ver string) StudentsController {
	return &studentsController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *studentsController) Init() {
	studentService := service.NewStudentsService(s.database)
	student := s.apiRoutes.Group("/v1/students")
	{
		student.POST("/create", studentService.Create)
		student.GET("/find-all", studentService.GetAllStudents)
	}
}