package store

import (
	"fmt"
	"os"

	"github.com/cisco/admin/gormdevice/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MySQLConnectionHelper() *gorm.DB {

	_ = godotenv.Load(".env")
	username := os.Getenv("username")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	host := os.Getenv("host")
	port := os.Getenv("port")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Device{})

	return db
}
