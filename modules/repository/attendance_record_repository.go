package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type AttendanceRecordRepository interface {
	Save(attendanceRecordRepository model.AttendanceRecord) (*model.AttendanceRecord, error)
}

type attendanceRecordRepository struct {
	conn *gorm.DB
}

func NewAttendanceRecordRepository(db *gorm.DB) AttendanceRecordRepository {
	return &attendanceRecordRepository{
		conn: db,
	}
}

func (a *attendanceRecordRepository) Save(attendanceRecordRepository model.AttendanceRecord) (*model.AttendanceRecord, error) {
	panic("unimplemented")
}