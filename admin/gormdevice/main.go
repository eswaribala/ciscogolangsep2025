package main

import (
	"fmt"

	"github.com/cisco/admin/gormdevice/store"
)

func main() {
	db := store.MySQLConnectionHelper()
	store.GetTableInstance(db)

	fmt.Println("Database connection established")
}
