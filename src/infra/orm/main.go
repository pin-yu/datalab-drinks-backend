package orm

import (
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/pinyu/datalab-drinks-backend/src/utils"
)

var dbSingleton *gorm.DB

var (
	basePath = utils.GetBasePath()
	dbPath   = filepath.Join(basePath, "../infra/database/dev.db")
)

// MigrateDB will migrate the database
func MigrateDB() {
	migrateOrder()
}

// DropDB will drop the database
func DropDB() {
	dropOrder()
}

func newDB() *gorm.DB {
	// if db has been instantiated, just return it (singleton)
	if dbSingleton != nil {
		return dbSingleton
	}

	// for the singleton db, declare error instead of using :=
	var err error
	dbSingleton, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	return dbSingleton
}
