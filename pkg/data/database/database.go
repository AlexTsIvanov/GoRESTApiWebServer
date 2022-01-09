package database

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connection to Database failed to open")
	}
	err = conn.AutoMigrate(&TypeItem{}, &MenuItem{}, &Status{}, &User{}, &Order{}, &OrderItem{})
	if err != nil {
		log.Println("Error migrating tables")
	}
	db = conn
}

func Connector() *gorm.DB {
	return db
}

func FromJSON(r io.Reader, m interface{}) error {
	e := json.NewDecoder(r)
	return e.Decode(m)
}
