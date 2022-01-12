package database

import (
	"../models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tuktuk.db"), &gorm.Config{})

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatalf("Unable to open SQLite database")
	}

	err = db.AutoMigrate(&models.File{}, &models.User{})

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatalf("Unable to auto migrate database changes")
	}

	return db
}
