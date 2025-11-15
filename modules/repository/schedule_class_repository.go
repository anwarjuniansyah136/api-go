package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type ScheduleClassRepository interface {
	Save(scheduleClass model.ScheduleClass) (*model.ScheduleClass, error)
	FindAll() (*[]model.ScheduleClass, error)
	FindById(id uint64) (*model.ScheduleClass, error)
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
	scheduleClass.UpdateAt = time.Now()
	if scheduleClass.ID == 0 {
		scheduleClass.CreateAt = time.Now()
	}

	if err := s.conn.Save(&scheduleClass).Error; err != nil {
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

func (s *scheduleClassRepository) FindById(id uint64) (*model.ScheduleClass, error) {
	var scheduleClass model.ScheduleClass

	err := s.conn.Where("id = ?", id).First(&scheduleClass).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &scheduleClass, nil
}