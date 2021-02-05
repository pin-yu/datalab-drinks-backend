package orm

import (
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/pinyu/datalab-drinks-backend/src/utils"
)

var singletonConnection *gorm.DB

var (
	basePath = utils.GetBasePath()
	dbPath   = filepath.Join(basePath, "../infrastructure/local/dev.db")
)

// MigrateDB will migrate the database
func MigrateDB() {
	migrateOrder()
}

// DropDB will drop the database
func DropDB() {
	dropOrder()
}

func newDBConnection() *gorm.DB {
	// if db has been instantiated, just return it (singleton)
	if singletonConnection != nil {
		return singletonConnection
	}

	// for the singleton db, declare error instead of using :=
	var err error
	singletonConnection, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	return singletonConnection
}
