package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	Save(teacher model.Teacher) (*model.Teacher, error)
	FindAll() (*[]model.Teacher, error)
}

type teacherRepository struct {
	conn *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{
		conn: db,
	}
}

func (t *teacherRepository) Save(teacher model.Teacher) (*model.Teacher, error) {
	teacher.CreateAt = time.Now()

	if err := t.conn.Create(&teacher).Error; err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (t *teacherRepository) FindAll() (*[]model.Teacher, error) {
	var teacher []model.Teacher
	err := t.conn.Find(&teacher).Error

	if err != nil {
		return nil, err
	}

	return  &teacher, err
}