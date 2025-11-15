package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoomService interface {
	Create(ctx *gin.Context)
	GetAllRoom(ctx *gin.Context)
}

type roomService struct {
	repository repository.RoomRepository
}

func NewRoomService(db *gorm.DB) RoomService {
	return &roomService{
		repository: repository.NewRoomRepository(db),
	}
}

func (r *roomService) Create(ctx *gin.Context) {
	var roomRequest model.RoomCreateRequest
	
	if err := ctx.ShouldBindJSON(&roomRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	room := model.Room{
		Name: roomRequest.Name,
		SchoolID: roomRequest.SchoolID,
	}

	result, err := r.repository.Save(room)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (r *roomService) GetAllRoom(ctx *gin.Context) {
	result, err := r.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}