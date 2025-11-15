package model

import "time"

type AttendanceSession struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	SchoolID  uint64    `json:"-"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	EndTime   time.Time `json:"end_time" gorm:"not null"`
	CreatedBy uint64    `json:"-"`
	CreatedAt time.Time `json:"-"`

	School School `json:"school" binding:"required" gorm:"foreignKey:SchoolID;references:ID"`
	User User `json:"user" binding:"required" gorm:"foreignKey:CreatedBy;references:ID"`
}

type AttendanceSessionCreateRequest struct {
	SchoolID uint64 `json:"school_id" binding:"required"`
	Name string `json:"name" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime string `json:"end_time" binding:"required"`
} 