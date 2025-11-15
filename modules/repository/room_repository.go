package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Save(room model.Room) (*model.Room, error)
	FindAll() (*[]model.Room, error)
	FindById(id uint64) (*model.Room, error)
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
	if room.ID == 0 {
		room.CreateAt = time.Now()
	}

	if err := r.conn.Save(&room).Error; err != nil {
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

func (r *roomRepository) FindById(id uint64) (*model.Room, error) {
	var room model.Room

	err := r.conn.Where("id = ?", id).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &room, nil
}