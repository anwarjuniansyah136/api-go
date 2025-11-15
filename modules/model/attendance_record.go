package model

import "time"

type AttendanceRecord struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement"`
	SessionID    uint64    `json:"-"`
	StudentID    uint64    `json:"-"`
	CheckinTime  time.Time `json:"checkin_time" gorm:" not null"`
	Latitude     float64   `json:"latitude,omitempty" gorm:"type:decimal(10,8)"`
	Longitude    float64   `json:"longitude,omitempty" gorm:"type:decimal(10,8)"`
	SelfieURL    string    `json:"selfie_url" gorm:"size:150;not null"`
	DistanceFrom float64   `json:"distance_from" gorm:"type:decimal(10,8)"`
	SchoolID     uint64    `json:"-"`
	Status       bool      `json:"status" gorm:"not null"`
	VerifiedAt   time.Time `json:"verified_at" gorm:"not null"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`

	AttendanceSession AttendanceSession `json:"attendance_session" required:"binding" gorm:"foreignKey:SessionID;references:ID"`
	Student Student `json:"student" binding:"required" gorm:"foreignKey:StudentID;references:ID"`
	School School `json:"school" binding:"required" gorm:"foreignKey:SchoolID;references:ID"`
}

type AttendanceRecordCreateRequest struct {
	SessionID uint64 `json:"session_id" binding:"required"`
	StudentID uint64 `json:"student_id" binding:"required"`
	CheckinTime string `json:"checkin_time" binding:"required"`
	Latitude float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	SelfieURL string `json:"selfie_url" binding:"required"`
	DistanceFrom float64 `json:"distance_from" binding:"required"`
	SchoolID uint64 `json:"school_id" binding:"required"`
	VerifiedAt string `json:"verified_at" binding:"required"`
}