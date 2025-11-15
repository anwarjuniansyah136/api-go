package model

import (
	"time"
)

type Teacher struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" binding:"required" gorm:"size:255;not null"`
	Address   string `json:"address" binding:"required" gorm:"size:255;not null"`
	Age       int `json:"age" binding:"required" gorm:"not null"`
	SubjectID uint64 `json:"-"`
	CreateAt  time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`

	Subject Subject `json:"subject" required:"binding" gorm:"foreignKey:SubjectID;references:ID"`
}

type TeacherCreateRequest struct{
	Name string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Age int `json:"age" binding:"required"`
	SubjectID uint64 `json:"subject_id" binding:"required"`
}