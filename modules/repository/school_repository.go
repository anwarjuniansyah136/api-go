package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type SchoolRepository interface {
	Save(school model.School) (*model.School, error)
	FindAll() (*[]model.School, error)
	FindById(id uint64) (*model.School, error)
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
	school.UpdateAt = time.Now()
	if school.ID == 0 {
		school.CreateAt = time.Now()
	}

	if err := s.conn.Save(&school).Error; err != nil {
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

func (s *schoolRepository) FindById(id uint64) (*model.School, error) {
	var school model.School

	err := s.conn.Where("id = ?", id).First(&school).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &school, nil
}