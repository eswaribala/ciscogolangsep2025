package models

import "time"

type Device struct {
	DeviceID    uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	HostName    string    `json:"host_name" gorm:"type:varchar(100);unique;not null"`
	Description string    `json:"description" gorm:"type:varchar(1024)"`
	IPAddress   string    `json:"ip_address" gorm:"type:varchar(45);not null;unique"`
	Location    string    `json:"location" gorm:"type:varchar(100);not null"`
	Status      bool      `json:"status" gorm:"type:boolean;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
