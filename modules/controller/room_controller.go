package controller

import (
	"api/modules/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoomController interface {
	Init()
}

type roomController struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewRoomController(apiR *gin.RouterGroup, db *gorm.DB, ver string) RoomController {
	return &roomController{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (r *roomController) Init() {
	roomService := service.NewRoleService(r.database)
	room := r.apiRoutes.Group("/v1/room")
	{
		room.POST("/create", roomService.Create)
		room.GET("/find-all", roomService.GetAllRole)
	}
}