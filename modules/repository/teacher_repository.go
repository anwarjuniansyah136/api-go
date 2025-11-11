package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	Save(teacher model.Teacher) (*model.Teacher, error)
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
	panic("unimplemented")
}