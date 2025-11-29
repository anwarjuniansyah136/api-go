package model

import "time"

type Role struct {
	ID       uint64    `gorm:"primaryKey;autoIncrement"`
	Name     string    `json:"name" gorm:"size:100;not null"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`

	User []User `gorm:"foreignKey:RoleID" json:"users,omitempty"`
}

type RoleCreateRequest struct {
	Name string `json:"name" binding:"required"`
}