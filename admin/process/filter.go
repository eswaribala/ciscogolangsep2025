package process

func FilterDevices(devices *[]string) *[]string {
	var filtered []string
	for _, device := range *devices {
		if device != "Client" && device != "Modem" {
			filtered = append(filtered, device)
		}
	}
	return &filtered
}
