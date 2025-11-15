package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoomService interface {
	Create(ctx *gin.Context)
	GetAllRoom(ctx *gin.Context)
	FindById(ctx *gin.Context)
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
			"error": err,
		})
		return
	}

	room := model.Room{
		Name:     roomRequest.Name,
		SchoolID: roomRequest.SchoolID,
	}

	result, err := r.repository.Save(room)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (r *roomService) GetAllRoom(ctx *gin.Context) {
	result, err := r.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (r *roomService) FindById(ctx *gin.Context) {
	var id = ctx.Param("id")

	value, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	result, err := r.repository.FindById(value)
	if err == nil {
		if result == nil{
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":"teacher not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : "something wrong in our server",
		})
	}

	ctx.JSON(http.StatusOK, result)
}