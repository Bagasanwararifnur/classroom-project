package db

import (
	"cls-project/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Thread{}, &models.Comment{}, &models.Status{}, &models.Class{})
}

func GetDB() *gorm.DB {
	return db
}
