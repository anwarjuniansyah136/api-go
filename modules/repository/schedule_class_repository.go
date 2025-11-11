package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type ScheduleClassRepository interface {
	Save(scheduleClass model.ScheduleClass) (*model.ScheduleClass, error)
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
	panic("unimplemented")
}