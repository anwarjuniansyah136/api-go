package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	Save(teacher model.Teacher) (*model.Teacher, error)
	FindAll() (*[]model.Teacher, error)
	FindByID(id uint64) (*model.Teacher, error)
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
	teacher.UpdateAt = time.Now()
	if teacher.ID == 0 {
		teacher.CreateAt = time.Now()
	}

	if err := t.conn.Save(&teacher).Error; err != nil {
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

	return &teacher, err
}

func (t *teacherRepository) FindByID(id uint64) (*model.Teacher, error) {
	var teacher model.Teacher
	err := t.conn.Where("id = ?", id).First(&teacher).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &teacher, nil
}
