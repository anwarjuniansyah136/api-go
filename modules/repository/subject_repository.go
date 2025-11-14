package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	Save(subject model.Subject) (*model.Subject, error)
	FindAll() (*[]model.Subject, error)
}

type subjectRepository struct {
	conn *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{
		conn: db,
	}
}

func (s *subjectRepository) Save(subject model.Subject) (*model.Subject, error) {
	subject.CreateAt = time.Now()

	if err := s.conn.Create(&subject).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

func (s *subjectRepository) FindAll() (*[]model.Subject, error) {
	var subjects []model.Subject

	err := s.conn.Find(&subjects).Error
	if err != nil {
		return nil, err
	}

	return &subjects, nil
}