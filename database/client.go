package database

import (
	"github.com/bandgren/classified-ads/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var Instance *gorm.DB
var err error

func Connect() {
	Instance, err = gorm.Open(mysql.Open(os.Getenv("SQL_CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Ad{})
	log.Println("Database Migration Completed...")
}
