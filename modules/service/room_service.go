package service

import (
	"api/modules/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoomService interface {
	Create(ctx *gin.Context)
}

type roomService struct {
	repository repository.RoomRepository
}

func NewRoomRepository(db *gorm.DB) RoomService {
	return &roomService{
		repository: repository.NewRoomRepository(db),
	}
}

func (r *roomService) Create(ctx *gin.Context) {
	panic("unimplemented")
}