package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pinyu/datalab-drinks-backend/src/utils"
)

var (
	basePath = utils.GetBasePath()
)

// GetDBPath return db path corresponding to the GIN_MODE
func GetDBPath() string {
	var dbMode string

	env := os.Getenv("GIN_MODE")
	switch env {
	case "release":
		dbMode = "release.db"
	case "test":
		dbMode = "test.db"
	default:
		dbMode = "dev.db"
	}

	return filepath.Join(basePath,
		fmt.Sprintf("../infrastructure/local/%v", dbMode))
}

// GetCamaYamlPath return cama yaml path
func GetCamaYamlPath() string {
	return filepath.Join(utils.GetBasePath(), "../../assets/cama_menu.yaml")
}
