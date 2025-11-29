package controller

import (
	"api/modules/middleware"
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeacherController interface {
	Init()
}

type teacherController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewTeacherController(apiR *gin.RouterGroup, db *gorm.DB, ver string) TeacherController {
	return &teacherController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (t *teacherController) Init() {
	teacherService := service.NewTeacherService(t.database)
	teacher := t.apiRoutes.Group("/v1/teacher", middleware.AuthMiddleWare())
	{
		teacher.POST("/create", teacherService.Create)
		teacher.GET("/find-all", teacherService.GetAllTeacher)
	}
}