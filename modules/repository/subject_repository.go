package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	Save(subject model.Subject) *model.Subject
}

type subjectRepository struct {
	conn *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{
		conn: db,
	}
}

func (s *subjectRepository) Save(subject model.Subject) *model.Subject {
	panic("unimplemented")
}