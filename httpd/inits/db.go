package inits

import (
	"github.com/fronomenal/go_jwt/httpd/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}

func Sync(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
