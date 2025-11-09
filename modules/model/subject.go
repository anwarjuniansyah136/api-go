package model

import "time"

type Subject struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	SchoolID    uint64    `json:"-"`
	SubjectCode string    `json:"subject_code" gorm:"size:100;not null"`
	SubjectName string    `json:"subject_name" gorm:"size:100;not null"`
	IsActive    bool      `json:"is_active" gorm:"type:boolean;default:true"`
	CreateAt    time.Time `json:"-"`
	UpdateAt    time.Time `json:"-"`

	School School `json:"school" binding:"required" gorm:"foreignKey:SchoolID;references:ID"`
}
