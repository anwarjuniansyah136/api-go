package model

import "time"

type AttendanceRecord struct {
	ID           uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
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
	Student Student `json:"student" binding:"required" gorm:"foreignKey:StudendID;references:ID"`
	School School `json:"school" binding:"required" gorm:"foreignKey:SchoolID;references:ID"`
}
