package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type DeviceLogsRepository interface {
	Save(device model.DeviceLog) (*model.DeviceLog, error)
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
	panic("unimplemented")
}