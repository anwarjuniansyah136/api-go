package controller

import (
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	Init()
}

type userController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewUserController(apiR *gin.RouterGroup, db *gorm.DB, ver string) UserController {
	return &userController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (u *userController) Init() {
	userService := service.NewUserService(u.database)
	user := u.apiRoutes.Group("/v1/user")
	{
		user.POST("/create", userService.Create)
		user.GET("/find-all", userService.GetAllUser)
		// user.PUT("/update/:id")
	}
}