package persistence

import (
	"log"
	"os"
	"os/exec"
)

func setupTestDB() {
	if os.Getenv("GIN_MODE") != "test" {
		log.Fatal("please set GIN_MODE=test")
	}

	NewOrderRepository().DropTable()
	NewOrderRepository().MigrateTable()
}

func cleanTestDB() {
	exec.Command("rm", DbPath())
}
