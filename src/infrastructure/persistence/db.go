package persistence

import (
	"log"

	"github.com/pin-yu/datalab-drinks-backend/src/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var singletonDriver *gorm.DB

var (
	dbPath = config.GetDBPath()
)

// DbPath returns current db path
func DbPath() string {
	return dbPath
}

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
