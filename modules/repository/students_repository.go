package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Save(student model.Student) (*model.Student, error)
	FindAll() (*[]model.Student, error)
	FindById(id uint64) (*model.Student, error)
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
	if student.ID == 0 {
		student.CreateAt = time.Now()
	}

	if err := s.conn.Save(&student).Error; err != nil {
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

func (s *studentRepository) FindById(id uint64) (*model.Student, error) {
	var student model.Student

	err := s.conn.Where("id = ?", id).First(&student).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &student, nil
}
