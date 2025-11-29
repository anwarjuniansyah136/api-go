package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type AttendanceRecordRepository interface {
	Save(attendanceRecord model.AttendanceRecord) (*model.AttendanceRecord, error)
	FindAll() (*[]model.AttendanceRecord, error)
	FindById(id uint64) (*model.AttendanceRecord, error)
}

type attendanceRecordRepository struct {
	conn *gorm.DB
}

func NewAttendanceRecordRepository(db *gorm.DB) AttendanceRecordRepository {
	return &attendanceRecordRepository{
		conn: db,
	}
}

func (a *attendanceRecordRepository) Save(attendanceRecord model.AttendanceRecord) (*model.AttendanceRecord, error) {
	attendanceRecord.UpdatedAt = time.Now()
	if attendanceRecord.ID == 0 {
		attendanceRecord.CreatedAt = time.Now()
	}

	if err := a.conn.Save(&attendanceRecord).Error; err != nil {
		return nil, err
	}

	return &attendanceRecord, nil
}

func (a *attendanceRecordRepository) FindAll() (*[]model.AttendanceRecord, error) {
	var attendanceRecords []model.AttendanceRecord

	err := a.conn.Find(&attendanceRecords).Error
	if err != nil {
		return nil, err
	}

	return &attendanceRecords, nil
}

func (a *attendanceRecordRepository) FindById(id uint64) (*model.AttendanceRecord, error) {
	var attendanceRecord model.AttendanceRecord

	err := a.conn.Where("id = ?", id).First(&attendanceRecord).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &attendanceRecord, nil
}
