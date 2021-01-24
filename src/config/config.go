package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GetDbDriver returns a database driver corresponding to the APP_ENV
func GetDbDriver() gorm.Dialector {
	env := os.Getenv("APP_ENV")
	var dialector gorm.Dialector

	if env == "production" {
		dialector = nil
	} else {
		dialector = sqlite.Open("dev.db")
	}

	return dialector
}
