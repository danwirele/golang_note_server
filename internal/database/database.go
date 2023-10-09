package database

import (
	"notes_server/internal/domain/note"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func DatabaseProvider() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	gopath := viper.GetString("db.path")
	database, err := gorm.Open(sqlite.Open(gopath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	database.AutoMigrate(note.Note{})

	db = database
	return db, nil
}
