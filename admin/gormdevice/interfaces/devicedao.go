package interfaces

import (
	"github.com/cisco/admin/gormdevice/store"
)

type DeviceDAO interface {
	GetAllDevices() ([]*store.Device, error)
	GetDeviceByID(id uint) (*store.Device, error)
	CreateDevice() (*store.Device, error)
	UpdateDevice(location string, status bool) (device *store.Device, err error)
	DeleteDevice(id uint) (bool, error)
}
