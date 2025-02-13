package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDb(path string, dst ...interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to setup db")
	}

	err = db.AutoMigrate(dst...)
	if err != nil {
		panic("failed to migrate db")
	}

	return db, err
}
