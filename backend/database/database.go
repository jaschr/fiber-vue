package database

import (
	"log"
	"os"

	"github.com/gahtoe/pfol/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

// Connect Database
func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.User{})

	DB = DbInstance{
		Db: db,
	}
}
