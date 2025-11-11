package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Save(room model.Room) (*model.Room, error)
}

type roomRepository struct {
	conn *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{
		conn: db,
	}
}

func (r *roomRepository) Save(room model.Room) (*model.Room, error) {
	panic("unimplemented")
}