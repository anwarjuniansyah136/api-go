package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Save(student model.Student) (*model.Student, error)
}

type studentRepository struct {
	conn *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{
		conn: db,
	}
}

func (s *studentRepository) Save(student model.Student) (*model.Student, error) {
	panic("unimplemented")
}
