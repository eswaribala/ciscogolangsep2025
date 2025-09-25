package store

type Device struct {
	DeviceID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	HostName    string `json:"host_name" gorm:"type:varchar(100);unique;not null"`
	Description string `json:"description" gorm:"type:varchar(1024)"`
	IPAddress   string `json:"ip_address" gorm:"type:varchar(45);not null;unique"`
	Location    string `json:"location" gorm:"type:varchar(100);not null"`
	Status      bool   `json:"status" gorm:"type:boolean;not null"`
	CreatedAt   string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   string `json:"updated_at" gorm:"autoUpdateTime"`
}

// method implementations
func (d *Device) CreateDevice() (*Device, error) {

	db := MySQLConnectionHelper()
	result := db.Create(d)
	if result.Error != nil {
		return nil, result.Error
	}
	return d, nil

}

func (d *Device) GetAllDevices() ([]*Device, error) {
	var devices []*Device
	db := MySQLConnectionHelper()
	result := db.Find(&devices)
	if result.Error != nil {
		return nil, result.Error
	}
	return devices, nil
}

func (d *Device) GetDeviceByID(id uint) (*Device, error) {
	var device Device
	db := MySQLConnectionHelper()
	result := db.First(&device, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &device, nil
}

func (d *Device) UpdateDevice(location string, status bool) (*Device, error) {
	db := MySQLConnectionHelper()
	result := db.Model(d).Updates(Device{Location: location, Status: status, UpdatedAt: "2025-09-21"})
	if result.Error != nil {
		return nil, result.Error
	}
	return d, nil
}

func (d *Device) DeleteDevice(id uint) (bool, error) {
	db := MySQLConnectionHelper()
	result := db.Delete(&Device{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
