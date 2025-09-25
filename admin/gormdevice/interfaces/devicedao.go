package interfaces

import (
	"github.com/cisco/admin/gormdevice/models"
)

type DeviceDAO interface {
	GetAllDevices() ([]*models.Device, error)
	GetDeviceByID(id uint) (*models.Device, error)
	CreateDevice() (*models.Device, error)
	UpdateDevice(location string, status bool) (device *models.Device, err error)
	DeleteDevice(id uint) (bool, error)
}
