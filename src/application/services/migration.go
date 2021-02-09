package services

import (
	"log"

	"github.com/pinyu/datalab-drinks-backend/src/infrastructure/persistence"
)

// MigrateTable is a special service for migrating database tables
func MigrateTable() {
	orderRepo := persistence.NewOrderRepository()
	orderRepo.MigrateTable()

	log.Println("database has been migrated")
}

// DropTable is a special service for drop database tables
func DropTable() {
	orderRepo := persistence.NewOrderRepository()
	orderRepo.DropTable()

	log.Println("database has been dropped")
}
