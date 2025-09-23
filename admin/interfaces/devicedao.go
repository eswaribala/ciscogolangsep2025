package interfaces

type DeviceDAO interface {
	//method
	Save() (bool, error)
	Update() (bool, error)
}
