package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type DeviceLogsRepository interface {
	Save(device model.DeviceLog) (*model.DeviceLog, error)
	FindAll() (*[]model.DeviceLog, error)
}

type deviceLogsRepository struct {
	conn *gorm.DB
}

func NewDeviceLogsRepository(db *gorm.DB) DeviceLogsRepository {
	return &deviceLogsRepository{
		conn: db,
	}
}

func (d *deviceLogsRepository) Save(device model.DeviceLog) (*model.DeviceLog, error) {
	device.CreateAt = time.Now()

	if err := d.conn.Create(&device).Error; err != nil {
		return nil, err
	}

	return &device, nil
}

func (d *deviceLogsRepository) FindAll() (*[]model.DeviceLog, error) {
	var deviceLogs []model.DeviceLog

	err := d.conn.Find(&deviceLogs).Error
	if err != nil {
		return nil, err
	}

	return &deviceLogs, nil
}