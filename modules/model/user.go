package model

import "time"

type User struct {
	ID       uint64    `gorm:"primaryKey;autoIncrement"`
	SchoolID uint64    `json:"-"`
	RoleID   uint64    `json:"-"`
	FullName string    `json:"full_name" gorm:"size:100"`
	Email    string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string    `json:"-" gorm:"type:text;not null"`
	Profile  string    `json:"profile" gorm:"size:255"`
	IsActive bool      `json:"is_active" gorm:"default:true" `
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`

	Role   Role   `json:"role" binding:"required" gorm:"foreignKey:RoleID;references:ID"`
	School School `json:"school" binding:"required" gorm:"foreignKey:SchoolID;references:ID"`
}