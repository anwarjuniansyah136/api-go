package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type ScheduleClassRepository interface {
	Save(scheduleClass model.ScheduleClass) (*model.ScheduleClass, error)
	FindAll() (*[]model.ScheduleClass, error)
}

type scheduleClassRepository struct {
	conn *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleClassRepository {
	return &scheduleClassRepository{
		conn: db,
	}
}

func (s *scheduleClassRepository) Save(scheduleClass model.ScheduleClass) (*model.ScheduleClass, error) {
	scheduleClass.CreateAt = time.Now()

	if err := s.conn.Create(&scheduleClass).Error; err != nil {
		return nil, err
	}

	return &scheduleClass, nil
}

func (s *scheduleClassRepository) FindAll() (*[]model.ScheduleClass, error) {
	var scheduleClasses []model.ScheduleClass

	err := s.conn.Find(&scheduleClasses).Error
	if err != nil {
		return nil, err
	}

	return &scheduleClasses, nil
}