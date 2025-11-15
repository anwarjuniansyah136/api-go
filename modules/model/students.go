package model

import "time"

type Student struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement"`
	UserID       uint64    `json:"-"`
	NISN         string    `json:"nisn" gorm:"type:varchar(20);unique;not null" `
	RoomID       uint64    `json:"-"`
	AcademicYear int       `json:"academic_year" gorm:"not null"`
	CreateAt     time.Time `json:"-"`

	User User `json:"user" binding:"required" gorm:"foreignKey:UserID;references:ID"`
	Room Room `json:"room" binding:"required" gorm:"foreignKey:RoomID;references:ID"`
}

type StudentCreateRequest struct {
	UserID uint64 `json:"user_id" binding:"required"`
	NISN string `json:"nisn" binding:"required"`
	RoomID uint64 `json:"room_id" binding:"required"`
	AcademicYear int `json:"academic_year" binding:"required"`
}