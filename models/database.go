package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

var conn *gorm.DB

func init() {
	dbLocation := os.Getenv("DB_PATH")
	if dbLocation == "" {
		dbLocation = "./database/database.db"
	}

	log.Printf("[DB] Attempting to connect")
	db, err := gorm.Open("sqlite3", dbLocation)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[DB] Connection Established")

	db.Debug().AutoMigrate(&Account{})

	conn = db
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return conn
}
