package persistence

import (
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/pinyu/datalab-drinks-backend/src/utils"
)

var singletonDriver *gorm.DB

var (
	basePath = utils.GetBasePath()
	dbPath   = filepath.Join(basePath, "../infrastructure/local/dev.db")
)

func newDBDriver() *gorm.DB {
	// if db has been instantiated, just return it (singleton)
	if singletonDriver != nil {
		return singletonDriver
	}

	// for the singleton db, declare error instead of using :=
	var err error
	singletonDriver, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	return singletonDriver
}
