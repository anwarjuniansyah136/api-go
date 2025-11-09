package model

import "time"

type Room struct {
	ID       uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string    `json:"name" gorm:"size:100;not null"`
	SchoolID uint64    `json:"-"`
	CreateAt time.Time `json:"-"`

	School School `json:"school" binding:"required" gorm:"foreignKey:SchoolID;references:ID"`
}