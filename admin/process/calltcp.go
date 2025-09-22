package process

import (
	"fmt"
	"net"
)

var addr *net.TCPAddr

// order 4
func init() {
	var err error
	addr, err = net.ResolveTCPAddr("tcp", "127.0.0.1:3306")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}
	fmt.Println("Resolved address from process package:", addr)
}

// order 5
func CallTCP() {
	fmt.Println("Main function...")
	fmt.Println("Connecting to server....")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server:", conn.RemoteAddr())
}
