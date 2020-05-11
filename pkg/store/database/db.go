package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func New() *gorm.DB {
	database, _ := gorm.Open("sqlite3", "test.db")

	return database
}

func Migrate(database *gorm.DB) {
	database.AutoMigrate(&Todo{})
	database.AutoMigrate(&User{})
}
