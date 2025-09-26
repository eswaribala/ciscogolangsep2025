package store

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once

func GenerateAutoMigration(db *gorm.DB) {
	println("Entering table creation")

	db.AutoMigrate(&Site{})
	println("Table Created")
}

func MySQLConnectionHelper() *gorm.DB {

	_ = godotenv.Load(".env")
	user, pass := VaultConnection()
	username := user
	password := pass
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
	//once.Do(func() { GenerateTable(db) })
	once.Do(func() { GenerateAutoMigration(db) })
}
