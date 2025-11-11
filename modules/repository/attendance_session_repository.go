package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type AttendanceSessionRepository interface {
	Save(attendanceSessionRepository model.AttendanceSession) (*model.AttendanceSession, error)
}

type attendanceSessionRepository struct {
	conn *gorm.DB
}

func NewAttendanceSessionRepository(db *gorm.DB) AttendanceSessionRepository {
	return &attendanceSessionRepository{
		conn: db,
	}
}

func (a *attendanceSessionRepository) Save(attendanceSessionRepository model.AttendanceSession) (*model.AttendanceSession, error) {
	panic("unimplemented")
}