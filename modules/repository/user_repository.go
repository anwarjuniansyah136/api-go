package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(teacher model.Teacher) *model.Teacher
}

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		conn: db,
	}
}

func (u *userRepository) Save(teacher model.Teacher) *model.Teacher {
	panic("unimplemented")
}