package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Save(room model.Room) (*model.Room, error)
	FindAll() (*[]model.Room, error)
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
	room.CreateAt = time.Now()

	if err := r.conn.Create(&room).Error; err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *roomRepository) FindAll() (*[]model.Room, error) {
	var rooms []model.Room

	err := r.conn.Find(&rooms).Error
	if err != nil {
		return nil, err
	}

	return &rooms, nil
}