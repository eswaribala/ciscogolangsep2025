package main

func main() {
	devices := []string{"router", "switch", "hub", "gateway"}
	ports := []int{22, 67, 78, 1068}

	for i, device := range devices {

		if ports[i] == 22 {
			println("SSH port detected!")
			continue
		}
		if ports[i] > 1027 {
			println("Beyond Reserved Port!")
			break
		}

		println("Device:", device, "Port:", ports[i])
	}

}
