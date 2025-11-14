package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Save(student model.Student) (*model.Student, error)
	FindAll() (*[]model.Student, error)
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
	student.CreateAt = time.Now()

	if err := s.conn.Create(&student).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *studentRepository) FindAll() (*[]model.Student, error) {
	var students []model.Student

	err := s.conn.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return &students, nil
}