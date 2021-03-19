package services

import (
	"log"
	"os"
	"os/exec"

	"github.com/pin-yu/datalab-drinks-backend/src/infrastructure/persistence"
)

// SetupTestDB will setup a test database
func SetupTestDB() {
	if os.Getenv("GIN_MODE") != "test" {
		log.Fatal("please set GIN_MODE=test")
	}

	DropTable()
	MigrateTable()
}

// CleanTestDB will clean the test database
func CleanTestDB() {
	DeleteDB()
}

// MigrateTable is a special service for migrating database tables
func MigrateTable() {
	orderRepo := persistence.NewOrderRepository()
	orderRepo.MigrateTable()

	log.Println("database has been migrated")
}

// DropTable is a special service for dropping database tables
func DropTable() {
	orderRepo := persistence.NewOrderRepository()
	orderRepo.DropTable()

	log.Println("database has been dropped")
}

// DeleteDB deletes .db files
func DeleteDB() {
	cmd := exec.Command("rm", persistence.DbPath())
	stdout, err := cmd.Output()

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(string(stdout))
}
