package model

import "time"

type ScheduleClass struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	SchoolID  uint64    `json:"-"`
	ClassName string    `json:"class_name" gorm:"size:100;not null"`
	SubjectID uint64    `json:"-"`
	TeacherID uint64    `json:"-"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	EndTime   time.Time `json:"end_time" gorm:"not null"`
	RoomID    uint64    `json:"-"`
	CreateAt  time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`

	School School `json:"school" binding:"required" gorm:"foreignKey:SchoolID;references:ID"`
	Subject Subject `json:"subject" binding:"required" gorm:"foreignKey:SubjectID;references:ID"`
	Teacher Teacher `json:"teacher" binding:"required" gorm:"foreignKey:TeacherID;references:ID"`
	Room Room `json:"room" binding:"required" gorm:"foreignKey:RoomID;references:ID"`
}

type ScheduleClassCreateRequest struct {
	SchoolID uint64 `json:"school_id" binding:"required"`
	ClassName string `json:"class_name" binding:"required"`
	SubjectID uint64 `json:"subject_id" binding:"required"`
	TeacherID uint64 `json:"steacher_id" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime string `json:"end_time" binding:"required"`
	RoomID uint64 `json:"role_id" binding:"required"`
}