package orm

import (
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/pinyu/datalab-drinks-backend/src/utils"
)

var db *gorm.DB

var (
	basePath = utils.GetBasePath()
	dbPath   = filepath.Join(basePath, "../infra/database/dev.db")
)

// NewDB will connect to the database and migrate the tables
func NewDB() *gorm.DB {
	log.Printf("db: %v\n", db)

	// if db has been instantiated, just return it (singleton)
	if db != nil {
		return db
	}

	// for the singleton db, declare error without using :=
	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	migrateOrder(db)

	return db
}
