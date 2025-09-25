package store

type Site struct {
	ID               uint   `gorm:"primaryKey"`
	Name             string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Location         string `gorm:"type:varchar(100);not null"`
	DeviceInterfaces []DeviceInterface `gorm:"foreignKey:SiteID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status           bool  `gorm:"not null;default:true"`
}
