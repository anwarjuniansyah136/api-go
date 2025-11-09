package model

import "time"

type Student struct {
	ID           uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID       uint64    `json:"-"`
	NISN         string    `json:"nisn" gorm:"type:varchar(20);unique;not null" `
	RoomID       uint64    `json:"-"`
	AcademicYear int       `json:"academic_year" gorm:"not null"`
	CreateAt     time.Time `json:"-"`

	User User `json:"user" binding:"required" gorm:"foreignKey:UserID;references:ID"`
	Room Room `json:"room" binding:"required" gorm:"foreignKey:RoomID;references:ID"`
}