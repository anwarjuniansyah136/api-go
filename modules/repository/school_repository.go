package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type SchoolRepository interface {
	Save(school model.School) (*model.School, error)
	FindAll() (*[]model.School, error)
}

type schoolRepository struct {
	conn *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &schoolRepository{
		conn: db,
	}
}

func (s *schoolRepository) Save(school model.School) (*model.School, error) {
	school.CreateAt = time.Now()
	school.UpdateAt = time.Now()

	if err := s.conn.Create(&school).Error; err != nil {
		return nil, err
	}
	return &school, nil
}

func (s *schoolRepository) FindAll() (*[]model.School, error) {
	var schools []model.School

	err := s.conn.Find(&schools).Error
	if err != nil {
		return nil, err
	}

	return &schools, nil
}