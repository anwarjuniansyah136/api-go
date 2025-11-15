package model

import "time"

type DeviceLog struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	UserID    uint64    `json:"-"`
	DeviceID  string    `json:"devide_id" gorm:"size:100;not null"`
	Platform  string    `json:"platform" gorm:"size:50"`
	IPAddress string    `json:"ip_address" gorm:"size:100;not null"`
	CreateAt  time.Time `json:"-"`

	User User `json:"user" binding:"required" gorm:"foreignKey:UserID;references:ID"`
}

type DeviceLogCreateRequest struct {
	UserID uint64 `json:"user_id" binding:"required"`
	DeviceID string `json:"device_id" binding:"reqired"`
	Platform string `json:"platform" binding:"required"`
	IPAddress string `json:"ip_address" binding:"required"`
}