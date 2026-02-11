package handlers

import (
	"log"

	"github.com/teddy-137/task_manager_api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(models.Task{})
}

func GetDB() *gorm.DB {
	return db
}
