package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	Save(subject model.Subject) (*model.Subject, error)
	FindAll() (*[]model.Subject, error)
	FindById(id uint64) (*model.Subject, error)
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
	subject.UpdateAt = time.Now()

	if subject.ID == 0 {
		subject.CreateAt = time.Now()
	}

	if err := s.conn.Save(&subject).Error; err != nil {
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

func (s *subjectRepository) FindById(id uint64) (*model.Subject, error) {
	var subject model.Subject

	err := s.conn.Where("id = ?", id).First(&subject).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &subject, nil
}