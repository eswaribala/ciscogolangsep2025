package interfaces

type DeviceDAO interface {
	//method
	Save() (bool, error)
	SaveToCSV(fileName string) (bool, error)
	Update() (bool, error)
}
