package store

import (
	"fmt"
	"os"
	"sync"

	"github.com/cisco/admin/gormdevice/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once

func GenerateTable(db *gorm.DB) {
	println("Entering table creation")

	db.AutoMigrate(&models.Device{})
	println("Table Created")
}

func MySQLConnectionHelper() *gorm.DB {

	_ = godotenv.Load(".env")
	username := os.Getenv("username")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	host := os.Getenv("host")
	port := os.Getenv("port")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetTableInstance(db *gorm.DB) {
	once.Do(func() { GenerateTable(db) })
}
