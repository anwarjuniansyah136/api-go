package controller

import (
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleController interface {
	Init()
}

type roleController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewRoleController(apiR *gin.RouterGroup, db *gorm.DB, ver string) RoleController {
	return &roleController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (r *roleController) Init() {
	roleService := service.NewRoleService(r.database)
	role := r.apiRoutes.Group("/v1/role")
	{
		role.POST("/create", roleService.Create)
		role.GET("/find-all", roleService.GetAllRole)
	}
}