package model

import "time"

type User struct {
	ID       uint64    `gorm:"primaryKey;autoIncrement"`
	SchoolID *uint64    `json:"-"`
	RoleID   *uint64    `json:"-"`
	FullName string    `json:"full_name" gorm:"size:100;not null"`
	Email    string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string    `json:"-" gorm:"type:text;not null"`
	Profile  string    `json:"profile" gorm:"size:255"`
	IsActive bool      `json:"is_active" gorm:"default:true"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`

	Role   Role   `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	School School `json:"school" gorm:"foreignKey:SchoolID;references:ID"`
}

type UserCreateRequest struct{
	FullName string `json:"full_name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateRequest struct{
	SchoolID *uint64    `json:"school_id"`
	RoleID   *uint64    `json:"school_id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Password *string    `json:"-"`
	Profile  string    `json:"profile"`
	IsActive bool      `json:"is_active"`
}