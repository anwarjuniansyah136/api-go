package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type DeviceLogsRepository interface {
	Save(device model.DeviceLog) (*model.DeviceLog, error)
	FindAll() (*[]model.DeviceLog, error)
	FindById(id uint64) (*model.DeviceLog, error)
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
	if device.ID == 0 {
		device.CreateAt = time.Now()
	}

	if err := d.conn.Save(&device).Error; err != nil {
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

func (d *deviceLogsRepository) FindById(id uint64) (*model.DeviceLog, error) {
	var device model.DeviceLog

	err := d.conn.Where("id = ?", id).First(&device).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &device, nil
}