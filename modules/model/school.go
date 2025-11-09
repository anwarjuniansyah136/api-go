package model

import "time"

type School struct {
	ID          uint64  `json:"id" gorm:"primaryKey;autoIncrement"`
	SchoolName  string  `json:"school_name" gorm:"size:255;not null"`
	Address     string  `json:"address" gorm:"type:text;not null"`
	Latitude    float64 `json:"latitude,omitempty" gorm:"type:decimal(10,8);not null"`
	Longitude   float64 `json:"longitude,omitempty" gorm:"type:decimal(10,8);not null"`
	RadiusMeter float64 `json:"radius_meter,omitempty" gorm:"type:decimal(10,8);not null"`
	CreateAt    time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}