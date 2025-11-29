package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type AttendanceSessionRepository interface {
	Save(attendanceSession model.AttendanceSession) (*model.AttendanceSession, error)
	FindAll() (*[]model.AttendanceSession, error)
	FindById(id uint64) (*model.AttendanceSession, error)
}

type attendanceSessionRepository struct {
	conn *gorm.DB
}

func NewAttendanceSessionRepository(db *gorm.DB) AttendanceSessionRepository {
	return &attendanceSessionRepository{
		conn: db,
	}
}

func (a *attendanceSessionRepository) Save(attendanceSession model.AttendanceSession) (*model.AttendanceSession, error) {
	attendanceSession.School.UpdateAt = time.Now()
	if attendanceSession.ID == 0 {
		attendanceSession.CreatedAt = time.Now()
	}

	if err := a.conn.Save(&attendanceSession).Error; err != nil {
		return nil, err
	}

	return &attendanceSession, nil
}

func (a *attendanceSessionRepository) FindAll() (*[]model.AttendanceSession, error) {
	var attendanceSessions []model.AttendanceSession

	err := a.conn.Find(&attendanceSessions).Error
	if err != nil {
		return nil, err
	}

	return &attendanceSessions, nil
}

func (a *attendanceSessionRepository) FindById(id uint64) (*model.AttendanceSession, error) {
	var attendanceSession model.AttendanceSession

	err := a.conn.Where("id = ?", id).First(&attendanceSession).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &attendanceSession, nil
}