package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type SchoolRepository interface {
	Save(school model.School) (*model.School, error)
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
	panic("unimplemented")
}