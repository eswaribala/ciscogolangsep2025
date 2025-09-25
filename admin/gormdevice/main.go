package main

import (
	"fmt"

	"github.com/cisco/admin/gormdevice/store"
)

func main() {
	store.MySQLConnectionHelper()

	fmt.Println("Database connection established")
}
