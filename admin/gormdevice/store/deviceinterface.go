package store

type DeviceInterface struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Description string `gorm:"type:varchar(255)"`
	DeviceType  string `gorm:"type:varchar(100);not null"`
	SiteID      uint   `gorm:"index"`
	Site        Site   `gorm:"foreignKey:SiteID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
